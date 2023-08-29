package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Btm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#504C4C"/><path fill="#FFF" d="m10.827 15.376l-1.745 3.698l.812 3.032l2.066.552l2.564 2.567l-.772.775l-6.112-1.64L6 18.248l1.23-1.228l5.342-5.343l-1.745 3.699l-1.745 3.698l5.442-5.442l-1.952-1.955l-1.745 3.699zm2.805 2.1l5.442 5.442l3.032-.812l.554-2.066l2.565-2.564l.775.772l-1.638 6.114L18.248 26l-6.569-6.572l1.953-1.952zm4.736-2.952l-5.442-5.442l-3.034.812l-.552 2.063l-2.565 2.567L6 13.752l1.64-6.114L13.752 6l1.228 1.23l5.34 5.34l-1.952 1.954zM20.04 9.34l-2.564-2.565l.772-.775l6.114 1.638L26 13.752l-6.572 6.569l-1.952-1.953l5.444-5.444l-.814-3.032l-2.066-.552z"/></g>`),
		g.Group(children),
	)
}