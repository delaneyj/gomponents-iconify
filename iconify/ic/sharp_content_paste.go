package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpContentPaste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2h-6.18C14.4.84 13.3 0 12 0S9.6.84 9.18 2H3v20h18V2zm-9 0c.55 0 1 .45 1 1s-.45 1-1 1s-1-.45-1-1s.45-1 1-1zm7 18H5V4h2v3h10V4h2v16z"/>`),
		g.Group(children),
	)
}