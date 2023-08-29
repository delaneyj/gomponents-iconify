package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 42v-8h6v8zm30 0v-8h6v8zM24 14v3m0-11v2m0 18v4M9 6v28M39 6v28m-15 4v4m-4-21H9v9l11-9Zm8 0h11v9l-11-9Z"/><path d="M18 15H9m21 0h9"/></g>`),
		g.Group(children),
	)
}