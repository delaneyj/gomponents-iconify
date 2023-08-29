package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.878 23.02l-9.9-9.9L11.1 3.016l9.9-.017v9.917L10.878 23.02Zm.001-2.827L19 12.085V5.002l-7.07.012l-8.122 8.108l7.071 7.07Zm4.117-11.19V7H17v2.004h-2.004Z"/>`),
		g.Group(children),
	)
}