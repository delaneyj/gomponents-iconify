package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInLineVerticalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 96L96 64h64Zm-32 96h64l-32-32Z" opacity=".2"/><path d="M224 128a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8ZM90.34 69.66A8 8 0 0 1 96 56h24V16a8 8 0 0 1 16 0v40h24a8 8 0 0 1 5.66 13.66l-32 32a8 8 0 0 1-11.32 0Zm25 2.34L128 84.69L140.69 72Zm50.35 114.34A8 8 0 0 1 160 200h-24v40a8 8 0 0 1-16 0v-40H96a8 8 0 0 1-5.66-13.66l32-32a8 8 0 0 1 11.32 0Zm-25-2.34L128 171.31L115.31 184Z"/></g>`),
		g.Group(children),
	)
}