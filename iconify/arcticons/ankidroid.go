package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ankidroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.63 9.21l.74 3.68l3.45 1.49l-3.29 1.84l-.34 3.78l-2.76-2.55l-3.67.83l1.57-3.41l-1.92-3.24l3.73.44l2.48-2.87Zm-13.07 15.7l4.46 3l5-2l-1.48 5.17L29 35.25l-5.37.19L20.69 40l-1.84-5l-5.19-1.34l4.23-3.31l-.33-5.36ZM37.5 4.5h-27a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}