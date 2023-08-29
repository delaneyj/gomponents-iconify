package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopMinusCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilLoopMinusCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M8.182 8.182a5.5 5.5 0 1 0 7.778 7.778a5.5 5.5 0 0 0-7.778-7.778Zm7.071 7.071A4.5 4.5 0 1 1 8.89 8.889a4.5 4.5 0 0 1 6.364 6.364Z" clip-rule="evenodd"/><path d="M15 17.121a1 1 0 0 1 1.414-1.414l3.789 3.789a1 1 0 0 1-1.414 1.414L15 17.12Zm-5.293-4.707a.5.5 0 1 1 0-1h4.485a.5.5 0 0 1 0 1H9.707Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilLoopMinusCircleFilled0)"/></g>`),
		g.Group(children),
	)
}