package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareMultipleTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.066 5H8.62a2.251 2.251 0 0 1 2.122-1.5h9.5a4.25 4.25 0 0 1 4.25 4.25v9.5c0 .976-.62 1.807-1.49 2.118v1.555a3.751 3.751 0 0 0 2.99-3.673v-9.5A5.75 5.75 0 0 0 20.241 2h-9.5a3.751 3.751 0 0 0-3.675 3ZM5.75 6A3.75 3.75 0 0 0 2 9.75v12.5A3.75 3.75 0 0 0 5.75 26h12.5A3.75 3.75 0 0 0 22 22.25V9.75A3.75 3.75 0 0 0 18.25 6H5.75Z"/>`),
		g.Group(children),
	)
}