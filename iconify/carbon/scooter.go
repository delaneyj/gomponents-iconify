package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scooter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 28a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2zM24 8h-6v2h6a1 1 0 0 1 0 2h-3a1 1 0 0 0-.98 1.196l.924 4.621L18.434 22h-2.69l-2.572-8.575A1.988 1.988 0 0 0 11.256 12H6v2h5.256l.6 2H7a5.006 5.006 0 0 0-5 5v2a1 1 0 0 0 1 1h1a4 4 0 0 0 8 0h7a1 1 0 0 0 .857-.485l3-5a1 1 0 0 0 .123-.711L22.22 14H24a3 3 0 0 0 0-6zM8 26a2.002 2.002 0 0 1-2-2h4a2.002 2.002 0 0 1-2 2zm-4-4v-1a3.003 3.003 0 0 1 3-3h5.456l1.2 4z"/>`),
		g.Group(children),
	)
}