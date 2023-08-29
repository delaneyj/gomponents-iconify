package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hulu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19.197 9.807H14.39a7.892 7.892 0 0 0-2.751.543V-.041H3.92v32.083h7.729V19.361a2.153 2.153 0 0 1 2.084-2.267h4.52c1.152 0 2.095.923 2.12 2.079V31.96h7.704V18.053c0-5.88-3-8.213-7.865-8.213z"/>`),
		g.Group(children),
	)
}