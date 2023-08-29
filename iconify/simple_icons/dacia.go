package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dacia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 8.646v2.23h8.252v2.248H0v2.23h9.112a.62.62 0 0 0 .489-.201L12 12.819l2.399 2.334a.62.62 0 0 0 .49.201H24v-2.23h-8.252v-2.248H24v-2.23h-9.112a.62.62 0 0 0-.489.201L12 11.181L9.601 8.847a.62.62 0 0 0-.49-.201Z"/>`),
		g.Group(children),
	)
}