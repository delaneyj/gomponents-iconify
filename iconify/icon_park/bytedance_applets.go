package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BytedanceApplets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M24 4V23L8 36"/><path d="M24 23L40 36"/><path d="M31 20L40 14"/><path d="M17 20L8 14"/><path d="M24 31V43"/></g>`),
		g.Group(children),
	)
}