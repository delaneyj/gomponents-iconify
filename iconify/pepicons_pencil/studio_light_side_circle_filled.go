package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioLightSideCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilStudioLightSideCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M17.168 7.872c.997-.665 2.332.05 2.332 1.248v3.263c0 1.198-1.335 1.913-2.332 1.248l-2.945-1.963l.554-.832l2.946 1.963a.5.5 0 0 0 .777-.416V9.12a.5.5 0 0 0-.777-.416l-2.946 1.964l-.554-.832l2.945-1.964Zm-.711 13.083a.5.5 0 0 1-.66.254L12 19.52v1.73a.5.5 0 0 1-1 0v-1.73l-3.797 1.688a.5.5 0 0 1-.406-.914L11 18.427v-4.175a.5.5 0 1 1 1 0v4.175l4.203 1.868a.5.5 0 0 1 .254.66Z"/><path d="M8 8.752a1.5 1.5 0 0 1 1.5-1.5h4a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5h-4a1.5 1.5 0 0 1-1.5-1.5v-4Zm1.5-.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4Z"/><path d="M14.146 8.605a.5.5 0 0 1 0-.707l3-3a.5.5 0 0 1 .708.707l-3 3a.5.5 0 0 1-.708 0Zm3.708 8a.5.5 0 0 1-.708 0l-3-3a.5.5 0 0 1 .707-.707l3 3a.5.5 0 0 1 0 .707Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilStudioLightSideCircleFilled0)"/></g>`),
		g.Group(children),
	)
}