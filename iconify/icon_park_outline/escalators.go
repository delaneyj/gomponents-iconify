package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Escalators(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35 13L13 35H4v9h9l22-22h9v-9h-9Zm-16 0l9-9m-6 0h6v6m-14 8l-9 9m6 0H5v-6"/>`),
		g.Group(children),
	)
}