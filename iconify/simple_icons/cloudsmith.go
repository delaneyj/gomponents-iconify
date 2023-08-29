package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudsmith(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 10.667v2.667L13.333 24h-2.666L0 13.334v-2.667L10.667 0h2.666L24 10.667Zm-12 6.869a5.535 5.535 0 1 0 0-11.07a5.535 5.535 0 0 0 0 11.07Z"/>`),
		g.Group(children),
	)
}