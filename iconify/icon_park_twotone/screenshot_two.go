package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTScreenshotTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24V4h40v20"/><path fill="#555" fill-rule="evenodd" d="M10 44a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z" clip-rule="evenodd"/><path d="M36 12C20.39 33.266 15.805 40.68 14.243 42.243a6 6 0 0 1-8.486 0"/><path fill="#555" fill-rule="evenodd" d="M38 44a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z" clip-rule="evenodd"/><path d="M42.243 42.243a6 6 0 0 1-8.486 0C32.195 40.68 27.61 33.266 12 12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTScreenshotTwo0)"/>`),
		g.Group(children),
	)
}