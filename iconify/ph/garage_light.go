package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 194h-10V98.67A14 14 0 0 0 223.77 87l-88-58.66a14 14 0 0 0-15.54 0L32.23 87A14 14 0 0 0 26 98.67V194H16a6 6 0 0 0 0 12h224a6 6 0 0 0 0-12ZM38 98.67a2 2 0 0 1 .89-1.67l88-58.67a2 2 0 0 1 2.22 0l88 58.67a2 2 0 0 1 .89 1.66V194h-28v-58a6 6 0 0 0-6-6H72a6 6 0 0 0-6 6v58H38ZM178 142v20h-44v-20Zm-56 20H78v-20h44Zm-44 12h44v20H78Zm56 0h44v20h-44Z"/>`),
		g.Group(children),
	)
}