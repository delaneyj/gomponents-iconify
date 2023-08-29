package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpAndDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 36v8h32V12h-8v8h-8v8h-8v8h-8Zm6-23l9-9m-6 0h6v6M10 27H4v-6m9-3l-9 9"/>`),
		g.Group(children),
	)
}