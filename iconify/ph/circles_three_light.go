package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesThreeLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M170 76a42 42 0 1 0-42 42a42 42 0 0 0 42-42Zm-42 30a30 30 0 1 1 30-30a30 30 0 0 1-30 30Zm60 24a42 42 0 1 0 42 42a42 42 0 0 0-42-42Zm0 72a30 30 0 1 1 30-30a30 30 0 0 1-30 30ZM68 130a42 42 0 1 0 42 42a42 42 0 0 0-42-42Zm0 72a30 30 0 1 1 30-30a30 30 0 0 1-30 30Z"/>`),
		g.Group(children),
	)
}