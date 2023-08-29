package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 19h9m9 0h-9m0 0v-6m0 0h6V5H6v8h6ZM9 9.01l.01-.011M12 9.01l.01-.011"/>`),
		g.Group(children),
	)
}