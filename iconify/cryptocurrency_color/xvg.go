package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xvg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#00CBFF"/><path fill="#FFF" d="M9.61 10.335L8 7h16l-1.592 3.335H24L15.951 27L8 10.335h1.61zm0 0l6.438 13.33l6.36-13.33H9.611z"/><path fill="#FFF" d="M16 24.5L8 7h15.999z" opacity=".504"/></g>`),
		g.Group(children),
	)
}