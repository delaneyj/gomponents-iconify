package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesFourBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M78 36a42 42 0 1 0 42 42a42 42 0 0 0-42-42Zm0 60a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm100 24a42 42 0 1 0-42-42a42 42 0 0 0 42 42Zm0-60a18 18 0 1 1-18 18a18 18 0 0 1 18-18ZM78 136a42 42 0 1 0 42 42a42 42 0 0 0-42-42Zm0 60a18 18 0 1 1 18-18a18 18 0 0 1-18 18Zm100-60a42 42 0 1 0 42 42a42 42 0 0 0-42-42Zm0 60a18 18 0 1 1 18-18a18 18 0 0 1-18 18Z"/>`),
		g.Group(children),
	)
}