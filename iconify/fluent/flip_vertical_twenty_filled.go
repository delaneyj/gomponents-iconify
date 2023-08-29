package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.66 2.122a.75.75 0 0 1 .34.628v5.5a.75.75 0 0 1-.75.75H2.75a.75.75 0 0 1-.302-1.436l12.5-5.5a.75.75 0 0 1 .712.058ZM6.317 7.5H14.5V3.9L6.317 7.5ZM16 17.5a.5.5 0 0 1-.71.454l-13-6A.5.5 0 0 1 2.5 11h13a.5.5 0 0 1 .5.5v6Z"/>`),
		g.Group(children),
	)
}