package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSimple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 80a32 32 0 1 0-32-32a32 32 0 0 0 32 32Zm0-48a16 16 0 1 1-16 16a16 16 0 0 1 16-16Zm102.86 100.12a8 8 0 0 1-11 2.74c-.35-.21-35.11-20.59-83.88-22.67V149l62 69.73a8 8 0 1 1-12 10.62L128 164l-58 65.31a8 8 0 1 1-12-10.62L120 149v-36.82c-49 2.08-83.52 22.46-83.88 22.68a8 8 0 1 1-8.23-13.72C29.6 120.11 70.45 96 128 96s98.4 24.11 100.12 25.14a8 8 0 0 1 2.74 10.98Z"/>`),
		g.Group(children),
	)
}