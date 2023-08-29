package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSomalia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#4189DD" d="M32 5H4a4 4 0 0 0-4 4v18a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V9a4 4 0 0 0-4-4z"/><path fill="#FFF" d="M19.405 16.066L18 11.741l-1.405 4.325h-4.548l3.679 2.673l-1.405 4.325L18 20.391l3.679 2.673l-1.405-4.325l3.679-2.673z"/>`),
		g.Group(children),
	)
}