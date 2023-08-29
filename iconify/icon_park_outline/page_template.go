package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageTemplate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M23 4H4v22h19V4Zm21 30H4v9h40v-9Zm0-30H31v8h13V4Zm0 14H31v8h13v-8Z"/>`),
		g.Group(children),
	)
}