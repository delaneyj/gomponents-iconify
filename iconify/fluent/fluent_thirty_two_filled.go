package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluentThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.447 2.106a1 1 0 0 0-.894 0l-8 4A1 1 0 0 0 8 7v17.5a1 1 0 0 0 .51.872l8 4.5A1 1 0 0 0 18 29v-9.382l7.447-3.724a1 1 0 0 0 0-1.788L19.237 11l6.21-3.106a1 1 0 0 0 0-1.788l-8-4Z"/>`),
		g.Group(children),
	)
}