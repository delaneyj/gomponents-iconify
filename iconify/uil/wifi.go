package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm0-8a7.06 7.06 0 0 0-4.95 2.05a1 1 0 0 0 0 1.41a1 1 0 0 0 1.41 0a5 5 0 0 1 7.08 0a1 1 0 0 0 .7.3a1 1 0 0 0 .76-1.71A7.06 7.06 0 0 0 12 11Zm0-4a11.08 11.08 0 0 0-7.78 3.22a1 1 0 0 0 1.42 1.42a9 9 0 0 1 12.72 0a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42A11.08 11.08 0 0 0 12 7Zm10.61.39a15 15 0 0 0-21.22 0a1 1 0 0 0 1.42 1.42a13 13 0 0 1 18.38 0a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		g.Group(children),
	)
}