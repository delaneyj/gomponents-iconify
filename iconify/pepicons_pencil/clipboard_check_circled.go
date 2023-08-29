package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheckCircled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.175 2.5a.5.5 0 0 1 .5-.5h6.643a.5.5 0 0 1 .5.5v3.875a.5.5 0 0 1-.5.5H6.675a.5.5 0 0 1-.5-.5V2.5Zm1 .5v2.875h5.643V3H7.175Z"/><path d="M4.5 17V5h2V4h-2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h7.854a4.02 4.02 0 0 1-.819-1H4.5Zm11-5.97c.35.045.685.133 1 .26V5a1 1 0 0 0-1-1h-2v1h2v6.03Z"/><path d="M15 18a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 1a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M16.582 13.882a.5.5 0 0 1 .078.703l-1.106 1.382a1 1 0 0 1-1.488.082l-.696-.695a.5.5 0 0 1 .708-.707l.696.695l1.105-1.382a.5.5 0 0 1 .703-.078Z"/></g>`),
		g.Group(children),
	)
}