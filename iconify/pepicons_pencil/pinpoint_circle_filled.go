package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinpointCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPinpointCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/><path d="M6 11.123C6 15.125 10.223 22 13 22s7-6.875 7-10.877C20 7.191 16.868 4 13 4s-7 3.191-7 7.123Zm13 0C19 14.643 15.096 21 13 21s-6-6.357-6-9.877C7 7.74 9.688 5 13 5s6 2.74 6 6.123Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPinpointCircleFilled0)"/></g>`),
		g.Group(children),
	)
}