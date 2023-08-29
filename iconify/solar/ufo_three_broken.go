package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UfoThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m3.636 20.363l4.121-4.12M10 22l2.707-2.222M2 14l2.222-2.707M7.856 4c1.15.09 2.579.636 4.09 1.56M5.275 6.583c.165 2.13 1.905 5.225 4.702 8.021c3.906 3.907 8.394 5.752 10.024 4.122c.826-.825.76-2.384-.014-4.229a13.974 13.974 0 0 0-.966-1.864m-4.802-2.271c-2.256-2.256-2.726-3.478-2.815-3.975c-.032-.174.052-.335.178-.46a4.73 4.73 0 0 1 6.688 0l.384.384a4.73 4.73 0 0 1 0 6.689c-.124.125-.286.21-.46.178c-.337-.06-1.008-.297-2.108-1.146"/><path fill="currentColor" d="M13.028 13.271a1 1 0 1 1-1.414-1.414a1 1 0 0 1 1.414 1.414ZM10.2 9.029a1 1 0 1 1-1.414-1.415A1 1 0 0 1 10.2 9.03Zm7.071 7.071a1 1 0 1 1-1.414-1.415A1 1 0 0 1 17.27 16.1Z"/></g>`),
		g.Group(children),
	)
}