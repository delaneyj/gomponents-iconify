package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChainCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilChainCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.346 14.902a3 3 0 0 1-.879-4.15l2.18-3.354a3 3 0 1 1 5.03 3.272l-2.18 3.353a3 3 0 0 1-4.15.88Zm-1.31.193a4.002 4.002 0 0 1-.407-4.889l2.18-3.353a4 4 0 1 1 6.707 4.362l-2.18 3.353a4 4 0 0 1-6.3.527Z"/><path d="M8.398 18.841a3 3 0 0 1-.879-4.15l2.18-3.353a3 3 0 1 1 5.03 3.271l-2.18 3.353a3 3 0 0 1-4.15.88Zm-1.31.194a4.002 4.002 0 0 1-.407-4.89l2.18-3.352a4 4 0 1 1 6.707 4.361l-2.18 3.353a4 4 0 0 1-6.3.528Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilChainCircleFilled0)"/></g>`),
		g.Group(children),
	)
}