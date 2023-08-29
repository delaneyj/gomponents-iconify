package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 6v2h7v18h2V8h7V6H4zm21 .625L21.5 11H24v10h-2.5l3.5 4.375L28.5 21H26V11h2.5L25 6.625z"/>`),
		g.Group(children),
	)
}