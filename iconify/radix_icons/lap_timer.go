package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LapTimer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5.5A.5.5 0 0 1 6 0h3a.5.5 0 0 1 0 1H8v1.12a6.363 6.363 0 0 1 2.992 1.016a.638.638 0 0 1 .066-.078l1-1a.625.625 0 0 1 .884.884l-.975.975A6.4 6.4 0 1 1 7 2.119V1H6a.5.5 0 0 1-.5-.5Zm-3.4 8a5.4 5.4 0 1 1 10.8 0a5.4 5.4 0 0 1-10.8 0Zm5.4 0V4.1a4.4 4.4 0 1 0 3.111 7.511L7.5 8.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}