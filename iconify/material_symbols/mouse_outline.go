package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.9 0-4.95-2.05T5 15V9q0-2.9 2.05-4.95T12 2q2.9 0 4.95 2.05T19 9v6q0 2.9-2.05 4.95T12 22Zm1-13h4q0-1.8-1.137-3.175T13 4.1V9ZM7 9h4V4.1q-1.725.35-2.863 1.725T7 9Zm5 11q2.075 0 3.538-1.463T17 15v-4H7v4q0 2.075 1.463 3.538T12 20Zm0-9Zm1-2Zm-2 0Zm1 2Z"/>`),
		g.Group(children),
	)
}