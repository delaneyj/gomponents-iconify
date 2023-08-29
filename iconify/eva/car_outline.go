package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCarOutline0"><g id="evaCarOutline1"><path id="evaCarOutline2" fill="currentColor" d="M21.6 11.22L17 7.52V5a1.91 1.91 0 0 0-1.81-2H3.79A1.91 1.91 0 0 0 2 5v10a2 2 0 0 0 1.2 1.88a3 3 0 1 0 5.6.12h6.36a3 3 0 1 0 5.64 0h.2a1 1 0 0 0 1-1v-4a1 1 0 0 0-.4-.78ZM20 12.48V15h-3v-4.92ZM7 18a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm5-3H4V5h11v10Zm7 3a1 1 0 1 1-1-1a1 1 0 0 1 1 1Z"/></g></g>`),
		g.Group(children),
	)
}