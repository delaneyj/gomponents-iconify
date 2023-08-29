package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M95.89 48a32 32 0 1 1 32 32a32 32 0 0 1-32-32Zm132.23 73.14C226.4 120.11 185.55 96 128 96s-98.4 24.11-100.12 25.14a8 8 0 0 0 8.24 13.72c.36-.22 34.91-20.6 83.88-22.68V149l-62 69.69a8 8 0 1 0 12 10.62L128 164l58 65.27a8 8 0 0 0 12-10.62L136 149v-36.81c48.77 2.08 83.53 22.46 83.88 22.67a8 8 0 1 0 8.24-13.72Z"/>`),
		g.Group(children),
	)
}