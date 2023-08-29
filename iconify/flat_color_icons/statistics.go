package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Statistics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#37474F"><path d="M23 5h2v36h-2z"/><path d="m25.817 32.772l1.414 1.414l-10.04 10.04l-1.414-1.414z"/><path d="m32.259 42.824l-1.414 1.414l-10.04-10.04l1.414-1.414z"/></g><path fill="#CFD8DC" d="M4 8h40v28H4z"/><g fill="#607D8B"><path d="M3 7h42v4H3zm0 28h42v2H3z"/><circle cx="31.5" cy="43.5" r="1.5"/><circle cx="16.5" cy="43.5" r="1.5"/></g><g fill="#C51162"><path d="m31.9 18.9l-5.9 6l-6-6l-8.1 8l2.2 2.2l5.9-6l6 6l8.1-8z"/><path d="m36 24l-7-7h7z"/></g>`),
		g.Group(children),
	)
}