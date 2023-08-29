package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Museum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M9.27 19h3.916l.602 18H8.749l.521-18zM6 38h38v2h3v3h2v2H1v-2h2v-3h3v-2zm40-24.188L25.002 5L4 13.812V15h42v-1.188zM8 16h34v2H8v-2zm28.736 3h3.914l.607 18h-5.046l.525-18zm-9.152 0h3.914l.6 18h-5.041l.527-18zm-9.154 0h3.915l.596 18h-5.039l.528-18z"/>`),
		g.Group(children),
	)
}