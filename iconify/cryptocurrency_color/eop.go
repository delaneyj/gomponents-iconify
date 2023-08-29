package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#35A7DF"/><path fill="#FEFFFE" d="M14.714 28L9.366 13.737l-2.16 9.84zM16.12 4.171l-5.863 7.132l5.863 14.983l5.897-14.983zM17.56 28l5.349-14.263l2.125 9.84z"/></g>`),
		g.Group(children),
	)
}