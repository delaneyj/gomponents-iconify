package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mygesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.17 12.825C27.72-1.271 13.896 2.597 13.388 11.389c-.488 14.058 15.781.225 16.022 12.101c-.036 4.503-.125 9.77-.088 14.252c-1.483 15.294-19.446.64-12.552-5.046c5.084-4.785 17.854-9.334 17.854-9.334"/>`),
		g.Group(children),
	)
}