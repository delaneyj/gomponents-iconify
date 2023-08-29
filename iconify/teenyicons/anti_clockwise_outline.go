package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntiClockwiseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m6.5 2.499l-.354-.354l-.353.354l.353.353L6.5 2.5Zm1-.5H7v1h.5v-1ZM2 8.495v-.5H1v.5h1ZM8.145.146l-1.999 2l.708.706L8.853.854L8.145.146ZM6.146 2.852l2 1.999l.707-.707l-2-1.999l-.707.707ZM7.5 3C10.537 3 13 5.461 13 8.496h1A6.499 6.499 0 0 0 7.5 2v1ZM13 8.495a5.499 5.499 0 0 1-5.5 5.496v1c3.589 0 6.5-2.909 6.5-6.496h-1ZM7.5 13.99A5.499 5.499 0 0 1 2 8.495H1a6.499 6.499 0 0 0 6.5 6.496v-1Z"/>`),
		g.Group(children),
	)
}