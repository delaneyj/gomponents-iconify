package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Location(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.832 2.688A4.056 4.056 0 0 0 8.02 1.5h-.04a4.056 4.056 0 0 0-4 4c-.013.75.198 1.487.606 2.117L7.734 14h.533l3.147-6.383c.409-.63.62-1.367.606-2.117a4.056 4.056 0 0 0-1.188-2.812zM7.925 2.5l.082.01l.074-.01a3.075 3.075 0 0 1 2.941 3.037a2.74 2.74 0 0 1-.467 1.568l-.02.034l-.017.035L8 12.279l-2.517-5.1l-.017-.039l-.02-.034a2.74 2.74 0 0 1-.467-1.568A3.074 3.074 0 0 1 7.924 2.5zm.612 2.169a1 1 0 1 0-1.112 1.663a1 1 0 0 0 1.112-1.663zM6.87 3.837a2 2 0 1 1 2.22 3.326a2 2 0 0 1-2.22-3.326z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}