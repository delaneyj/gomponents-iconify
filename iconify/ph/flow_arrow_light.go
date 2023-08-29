package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowArrowLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m244.24 75.76l-32-32a6 6 0 0 0-8.48 8.48L225.51 74H192a72 72 0 0 0-27.51 6.55c-13.18 6.08-29.84 19.08-34.41 46.46c-5.84 35-33 41.57-44.6 42.75a38 38 0 1 0 .08 12a69.61 69.61 0 0 0 21.76-5.57c13.17-5.86 29.86-18.8 34.6-47.21c7-42.26 48.34-43 50.08-43h33.51l-21.75 21.76a6 6 0 1 0 8.48 8.48l32-32a6 6 0 0 0 0-8.46ZM48 202a26 26 0 1 1 26-26a26 26 0 0 1-26 26Z"/>`),
		g.Group(children),
	)
}