package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaleCondomOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="m31.262 9.99l.733-2.85L27.886 6L7.09 18.517L6 22.787l2.733.764l-.732 2.85l2.734.763l-.733 2.85l2.734.764l-.733 2.85l2.734.763l-.733 2.85l2.734.764l-.733 2.85l4.1 1.145l20.788-12.512L42 25.208l-2.733-.764l.732-2.85l-2.734-.763l.733-2.85l-2.734-.764l.733-2.85l-2.734-.763l.733-2.85l-2.734-.764Zm.31 2.163l-2.734-.764l.731-2.846l-1.389-.386l-19.355 11.65l-.403 1.58l2.736.765l-.732 2.85l2.733.763l-.732 2.85l2.733.764l-.732 2.85l2.733.763l-.732 2.85l2.733.764l-.732 2.85l1.383.386l19.348-11.646l.411-1.59l-2.73-.763l.732-2.85l-2.733-.763l.732-2.85l-2.733-.764l.732-2.85l-2.733-.763l.732-2.85Z"/><path d="M24 31a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0 2a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z"/><path d="M24 27a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 2a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g>`),
		g.Group(children),
	)
}