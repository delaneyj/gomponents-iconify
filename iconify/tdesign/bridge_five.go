package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 3h24v18.001h-6v-7.603a8.417 8.417 0 0 0-1.346-1.053A8.644 8.644 0 0 0 12 11c-2.02 0-3.586.671-4.654 1.345A8.42 8.42 0 0 0 6 13.398V21H0V3Zm6 7.836A10.644 10.644 0 0 1 12 9a10.644 10.644 0 0 1 6 1.836V5.001h-.625L12 5H6v5.836Zm14-5.835v14h2v-14h-2ZM4 5H2v14h2V5Z"/>`),
		g.Group(children),
	)
}