package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NintendoSwitchSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M3.5 4a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M7 .5v12a.5.5 0 0 1-.5.5h-3A3.5 3.5 0 0 1 0 9.5v-6A3.5 3.5 0 0 1 3.5 0h3a.5.5 0 0 1 .5.5Zm-5 4a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/><path fill="currentColor" d="M11.5 10a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M15 5.5v6a3.5 3.5 0 0 1-3.5 3.5h-3a.5.5 0 0 1-.5-.5v-12a.5.5 0 0 1 .5-.5h3A3.5 3.5 0 0 1 15 5.5Zm-5 5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}