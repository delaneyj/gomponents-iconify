package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.281 6.281L11 23.563L3.719 16.28L2.28 17.72l8 8l.719.687l.719-.687l18-18z"/>`),
		g.Group(children),
	)
}