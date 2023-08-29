package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalFullDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M200 40v160a8 8 0 0 1-8 8H32a8 8 0 0 1-5.66-13.66l160-160A8 8 0 0 1 200 40Z" opacity=".2"/><path d="M198.12 25.23a16 16 0 0 0-17.44 3.46l-160 160A16 16 0 0 0 32 216h160a16 16 0 0 0 16-16V40a15.94 15.94 0 0 0-9.88-14.77ZM192 200H32L192 40Z"/></g>`),
		g.Group(children),
	)
}