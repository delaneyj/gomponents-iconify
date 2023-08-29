package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Roller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.5 15.938a.501.501 0 0 1-.5-.5v-3.906c0-.275.225-.5.5-.5H8V8.389c0-.248.179-.456.425-.494L16 5.93V3.032h-.5a.501.501 0 0 1-.5-.5c0-.275.225-.5.5-.5h1c.275 0 .5.225.5.5v3.857a.496.496 0 0 1-.423.494L9 8.849v2.183h.5c.275 0 .5.225.5.5v3.906c0 .275-.225.5-.5.5h-2z"/><path d="M14 3V2a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2z"/></g>`),
		g.Group(children),
	)
}