package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="368" cy="256" r="128" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32"/><rect width="480" height="256" x="16" y="128" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="128" ry="128"/>`),
		g.Group(children),
	)
}