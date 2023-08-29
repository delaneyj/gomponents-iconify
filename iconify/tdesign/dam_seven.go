package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DamSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.22 1h11.56l1 4H21v15h1v2H2v-2h1V5h2.22l1-4Zm1.06 4h9.44l-.5-2H7.78l-.5 2ZM5 20h3v-8h8v8h3V7H5v13Zm9 0v-6h-4v6h4Z"/>`),
		g.Group(children),
	)
}