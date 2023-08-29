package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1zM3 14a9 9 0 0 1 14.618-7.032l.675-.675a1 1 0 1 1 1.414 1.414l-.675.675A9 9 0 1 1 3 14zm10-4a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}