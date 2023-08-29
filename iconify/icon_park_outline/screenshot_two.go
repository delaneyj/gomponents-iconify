package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24V4h40v20"/><path d="M10 44a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z" clip-rule="evenodd"/><path d="M36 12C20.39 33.266 15.805 40.68 14.243 42.243a6 6 0 0 1-8.486 0"/><path d="M38 44a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z" clip-rule="evenodd"/><path d="M42.243 42.243a6 6 0 0 1-8.486 0C32.195 40.68 27.61 33.266 12 12"/></g>`),
		g.Group(children),
	)
}