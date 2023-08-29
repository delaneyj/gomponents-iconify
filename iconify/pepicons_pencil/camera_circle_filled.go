package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCameraCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M14 7.5H8A2.5 2.5 0 0 0 5.5 10v5A2.5 2.5 0 0 0 8 17.5h6a2.5 2.5 0 0 0 2.5-2.5v-5A2.5 2.5 0 0 0 14 7.5ZM6.5 10A1.5 1.5 0 0 1 8 8.5h6a1.5 1.5 0 0 1 1.5 1.5v5a1.5 1.5 0 0 1-1.5 1.5H8A1.5 1.5 0 0 1 6.5 15v-5Z"/><path d="m18.727 8.58l-2.976 1.936a.5.5 0 0 0-.228.414l-.027 2.612a.5.5 0 0 0 .227.425l3.004 1.952a.5.5 0 0 0 .773-.419V9a.5.5 0 0 0-.773-.42Zm-.227 6l-2.001-1.301l.021-2.07l1.98-1.287v4.658Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCameraCircleFilled0)"/></g>`),
		g.Group(children),
	)
}