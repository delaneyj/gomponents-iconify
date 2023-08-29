package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Female(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19.72 4.28l-1.44 1.44L21.563 9l-2.968 2.97A8.935 8.935 0 0 0 13 10c-4.96 0-9 4.04-9 9s4.04 9 9 9s9-4.04 9-9a8.94 8.94 0 0 0-1.97-5.594l2.97-2.97l3.28 3.283l1.44-1.44L24.437 9l3.28-3.28l-1.437-1.44L23 7.563l-3.28-3.28zM13 12c3.878 0 7 3.122 7 7s-3.122 7-7 7s-7-3.122-7-7s3.122-7 7-7z"/>`),
		g.Group(children),
	)
}