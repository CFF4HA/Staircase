package chrome

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/CFF4HA/Staircase/internal/engine/core"
	"github.com/CFF4HA/Staircase/internal/types"
	"github.com/chromedp/chromedp"
	"text/template"
)

type Client struct {

	// the following are private variables for our
	// scraping client.
	tmpl template.Template
}

// This function will be used to create a new client using sane defualts.
func NewClient() (*Client, error) {
	c := &Client{}

	return c, c.setupTemplate(engine)
}

// This function will set up the template for the javascript engine that
// is injected into the browser context.
func (c *Client) setupTemplate(t string) error {
	tmpl, err := template.New("engine").Parse(t)
	if err != nil {
		return fmt.Errorf("failed to parse template, err=%v", err)
	}

	c.tmpl = *tmpl
	core.Logger.Info("successfully set up template for chrome client")
	return nil
}

func (c *Client) QueryToJS(q types.RetrievalQuery) (string, error) {
	var injectedJavascript strings.Builder
	if err := c.tmpl.Execute(&injectedJavascript, q); err != nil {
		core.Logger.Error("failed to execute template for injected javascript", "err", err)
		return "", err
	}

	return injectedJavascript.String(), nil
}

// The Perform function is the primary entry point to the Client function.
// It is designed to be stateless and execute a series of jobs based on a
// given context.
func (c *Client) Perform(ctx context.Context, job types.RetrievalQuery) (*types.StaircaseEngineLog, error) {
	if err := chromedp.Run(ctx, chromedp.EmulateViewport(1920, 1080)); err != nil {
		panic(err)
	}
	if job.Site != "" {
		if err := chromedp.Run(ctx, chromedp.Navigate(job.Site)); err != nil {
			core.Logger.Error("failed to visit site", "site", job.Site, "err", err)
			return nil, err
		}

		if err := chromedp.Run(ctx, chromedp.Evaluate(regex, nil)); err != nil {
			core.Logger.Debug("failed to inject regex function", "err", err)
		}
	}

	injectedJavascript, err := c.QueryToJS(job)
	if err != nil {
		core.Logger.Error("failed to execute template for injected javascript", "err", err)
		return nil, err
	}

	log := &types.StaircaseEngineLog{}
	if err := chromedp.Run(
		ctx,
		chromedp.Evaluate(injectedJavascript, log),
	); err != nil {
		core.Logger.Error("failed to execute injected javascript", "err", err, "injected_script", injectedJavascript)

		var screenshot []byte
		chromedp.Run(
			ctx,
			chromedp.FullScreenshot(&screenshot, 90),
		)
		os.WriteFile("/tmp/screenshot.png", screenshot, 0644)

		return nil, err
	}

	return log, nil
}
