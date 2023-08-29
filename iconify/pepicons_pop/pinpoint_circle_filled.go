package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinpointCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPinpointCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 14.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0-5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z"/><path d="M5.5 11.123C5.5 15.366 9.882 22.5 13 22.5c3.118 0 7.5-7.134 7.5-11.377C20.5 6.917 17.146 3.5 13 3.5s-7.5 3.417-7.5 7.623Zm13 0c0 3.28-3.745 9.377-5.5 9.377s-5.5-6.097-5.5-9.377C7.5 8.013 9.967 5.5 13 5.5s5.5 2.513 5.5 5.623Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPinpointCircleFilled0)"/></g>`),
		g.Group(children),
	)
}