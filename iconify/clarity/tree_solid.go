package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2C10.8 1.7 4.8 7.3 4.5 14.5C4.2 21.7 9.8 27.7 17 28v-5.2l-5.2-5.2c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0L17 20v-6.2l-2.7-2.7c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0L19 13v3l3.3-3.3c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4L19 18.8V28c7.2-.3 12.8-6.3 12.5-13.5S25.2 1.7 18 2z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M18 28h-1v5c0 .6.4 1 1 1s1-.4 1-1v-5h-1z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}