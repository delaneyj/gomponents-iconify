package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineCameraIndoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 13v-1c0-.55-.45-1-1-1H9c-.55 0-1 .45-1 1v4c0 .55.45 1 1 1h4c.55 0 1-.45 1-1v-1l2 1.06v-4.12L14 13zm-2-7.5l6 4.5v9H6v-9l6-4.5M12 3L4 9v12h16V9l-8-6z"/>`),
		g.Group(children),
	)
}