package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyVirginMedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 23.95c0 9.496-9.368 9.337-15.29 0c5.492-9.505 15.29-9.494 15.29 0Zm-39 0c0-14.75 14.527-14.503 23.71 0c-8.515 14.767-23.71 14.752-23.71 0Z"/>`),
		g.Group(children),
	)
}