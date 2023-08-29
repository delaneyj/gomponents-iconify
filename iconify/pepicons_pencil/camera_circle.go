package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M14 7.5H8A2.5 2.5 0 0 0 5.5 10v5A2.5 2.5 0 0 0 8 17.5h6a2.5 2.5 0 0 0 2.5-2.5v-5A2.5 2.5 0 0 0 14 7.5ZM6.5 10A1.5 1.5 0 0 1 8 8.5h6a1.5 1.5 0 0 1 1.5 1.5v5a1.5 1.5 0 0 1-1.5 1.5H8A1.5 1.5 0 0 1 6.5 15v-5Z"/><path d="m18.727 8.58l-2.976 1.936a.5.5 0 0 0-.228.414l-.027 2.612a.5.5 0 0 0 .227.425l3.004 1.952a.5.5 0 0 0 .773-.419V9a.5.5 0 0 0-.773-.42Zm-.227 6l-2.001-1.301l.021-2.07l1.98-1.287v4.658Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}