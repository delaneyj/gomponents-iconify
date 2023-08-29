package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hockey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.938 15.969H2.032a1 1 0 0 1-1-1c0-.553.447-.916 1-.916h3.807c.062-.113.143-.305.178-.388L11.022.818a1.002 1.002 0 0 1 1.317-.517c.506.222.737.811.517 1.317l-5 12.75c-.251.592-.675 1.601-1.918 1.601zM10 15h2.938v1H10z"/>`),
		g.Group(children),
	)
}