package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19.03 4.28l-11 11l-.686.72l.687.72l11 11l1.44-1.44L10.187 16l10.28-10.28l-1.437-1.44z"/>`),
		g.Group(children),
	)
}