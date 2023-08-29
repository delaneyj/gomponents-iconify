package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularNodataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.6 20.8q-.275.275-.688.288T16.2 20.8q-.275-.275-.275-.7t.275-.7l1.4-1.4l-1.4-1.4q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.4 1.4l1.4-1.4q.275-.275.687-.287t.713.287q.275.275.275.7t-.275.7L20.425 18l1.375 1.4q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L19 19.425L17.6 20.8ZM4.425 22q-.675 0-.937-.613T3.7 20.3L20.275 3.725q.475-.475 1.088-.212t.612.937v8.25q-.65-.4-1.425-.588T19 11.925q-2.525 0-4.3 1.775T12.925 18q0 1.15.375 2.138T14.425 22h-10Z"/>`),
		g.Group(children),
	)
}