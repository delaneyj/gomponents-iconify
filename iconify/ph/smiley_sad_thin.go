package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileySadThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92ZM84 108a8 8 0 1 1 8 8a8 8 0 0 1-8-8Zm88 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm-.54 66a4 4 0 0 1-1.46 5.46a3.93 3.93 0 0 1-2 .54a4 4 0 0 1-3.46-2c-8.21-14.19-21.19-22-36.54-22s-28.33 7.81-36.54 22a4 4 0 0 1-6.92-4c9.55-16.52 25.4-26 43.46-26s33.91 9.48 43.46 26Z"/>`),
		g.Group(children),
	)
}