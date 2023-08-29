package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrangeSliceThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 84H8a4 4 0 0 0-4 4a124 124 0 0 0 248 0a4 4 0 0 0-4-4ZM71.53 150.13L124 97.66v74.24a83.67 83.67 0 0 1-52.47-21.77Zm-5.66-5.66A83.67 83.67 0 0 1 44.1 92h74.24ZM132 97.66l52.47 52.47A83.67 83.67 0 0 1 132 171.9Zm58.13 46.81L137.66 92h74.24a83.67 83.67 0 0 1-21.77 52.47ZM128 204A116.14 116.14 0 0 1 12.07 92h24a92 92 0 0 0 183.82 0h24A116.14 116.14 0 0 1 128 204Z"/>`),
		g.Group(children),
	)
}