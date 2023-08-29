package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5zM5.16 13.67A2.84 2.84 0 0 1 8 11.51a2.82 2.82 0 0 1 2.84 2.16a7.24 7.24 0 0 1-5.68 0zm6.84-.61a4.09 4.09 0 0 0-4-2.77a4.09 4.09 0 0 0-3.95 2.78A6.14 6.14 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.11 6.11 0 0 1 12 13.06z"/><path fill="currentColor" d="M8.05 4.3A2.33 2.33 0 0 0 5.8 6.7a2.33 2.33 0 0 0 2.25 2.4a2.33 2.33 0 0 0 2.25-2.4a2.33 2.33 0 0 0-2.25-2.4zm0 3.55a1.08 1.08 0 0 1-1-1.15a1.08 1.08 0 0 1 1-1.15a1.08 1.08 0 0 1 1 1.15a1.08 1.08 0 0 1-1 1.15z"/>`),
		g.Group(children),
	)
}