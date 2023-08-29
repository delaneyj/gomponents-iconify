package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bootstrapping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22h18v2H3zm9-10c-12 0-9 9-9 9h18s3-9-9-9Z"/><path fill="currentColor" d="m3 0l1 13c1.22-1.6 3.46-2.51 8-2.51s6.81.93 8 2.53L21 0Zm4 3h10v2H7Zm11 6H6V7h12Z"/>`),
		g.Group(children),
	)
}