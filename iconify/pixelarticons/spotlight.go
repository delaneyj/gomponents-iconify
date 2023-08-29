package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spotlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2h16v20H3V2h2zm14 18V4H5v16h14zM13 6H7v2h6V6zm-6 4h10v8H7v-8z"/>`),
		g.Group(children),
	)
}