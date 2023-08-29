package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotionLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 42h-48a6 6 0 0 0 0 12h18v130.64L109.26 45.11A6 6 0 0 0 104 42H40a6 6 0 0 0 0 12h18v148H40a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12H70V71.36l76.74 139.53A6 6 0 0 0 152 214h40a6 6 0 0 0 6-6V54h18a6 6 0 0 0 0-12ZM74.15 54h26.3l81.4 148h-26.3Z"/>`),
		g.Group(children),
	)
}