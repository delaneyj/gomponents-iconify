package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContractCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilContractCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M8.354 18.354a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Z"/><path d="M12.5 18a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 1 0v4Z"/><path d="M8 14.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1H8Zm6.354-2.146a.5.5 0 0 1-.708-.708l4-4a.5.5 0 0 1 .708.708l-4 4Z"/><path d="M14 12.5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1h-4Z"/><path d="M14.5 12a.5.5 0 0 1-1 0V8a.5.5 0 0 1 1 0v4Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilContractCircleFilled0)"/></g>`),
		g.Group(children),
	)
}