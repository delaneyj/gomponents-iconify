package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132 156.29V88a4 4 0 0 0-8 0v68.29a28 28 0 1 0 8 0ZM128 204a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm36-68V48a36 36 0 0 0-72 0v88a60 60 0 1 0 72 0Zm-36 100a52 52 0 0 1-29.71-94.68A4 4 0 0 0 100 138V48a28 28 0 0 1 56 0v90a4 4 0 0 0 1.71 3.28A52 52 0 0 1 128 236Z"/>`),
		g.Group(children),
	)
}