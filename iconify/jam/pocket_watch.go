package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PocketWatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 13h2a1 1 0 0 1 0 2H8a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v3zM7 5.732V5a1 1 0 1 1 2 0v.732a2 2 0 1 0-2 0zm-2.041.866a4 4 0 1 1 6.082 0A8.002 8.002 0 0 1 8 22A8 8 0 0 1 4.959 6.598zM8 20A6 6 0 1 0 8 8a6 6 0 0 0 0 12z"/>`),
		g.Group(children),
	)
}