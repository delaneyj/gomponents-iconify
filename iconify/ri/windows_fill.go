package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.001 5.479l7.377-1.016v7.127H3V5.48Zm0 13.042l7.377 1.017v-7.04H3v6.023Zm8.188 1.125L21.001 21v-8.502h-9.812v7.148Zm0-15.292v7.236h9.812V3l-9.812 1.354Z"/>`),
		g.Group(children),
	)
}