package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandApple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7c-3 0-4 3-4 5.5c0 3 2 7.5 4 7.5c1.088-.046 1.679-.5 3-.5c1.312 0 1.5.5 3 .5s4-3 4-5c-.028-.01-2.472-.403-2.5-3c-.019-2.17 2.416-2.954 2.5-3c-1.023-1.492-2.951-1.963-3.5-2c-1.433-.111-2.83 1-3.5 1c-.68 0-1.9-1-3-1zm3-3a2 2 0 0 0 2-2a2 2 0 0 0-2 2"/>`),
		g.Group(children),
	)
}