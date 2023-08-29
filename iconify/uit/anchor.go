package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 13.5h-3a.5.5 0 0 0 0 1h2.482a7 7 0 0 1-6.482 6.482V11h2a.5.5 0 0 0 0-1h-2V7.95A2.998 2.998 0 0 0 15 5a3 3 0 1 0-6 0a2.994 2.994 0 0 0 2.5 2.95V10h-2a.5.5 0 0 0 0 1h2v9.974c-3.419-.241-6.23-2.956-6.482-6.474H7.5a.5.5 0 0 0 0-1h-3a.5.5 0 0 0-.5.5a8.01 8.01 0 0 0 8 8a8.01 8.01 0 0 0 8-8a.5.5 0 0 0-.5-.5zM10 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0z"/>`),
		g.Group(children),
	)
}