package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoneAllRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 17.3l-4.25-4.25q-.3-.3-.288-.7t.313-.7q.3-.275.7-.288t.7.288l3.55 3.55l1.4 1.4l-.725.7q-.3.275-.7.288T6 17.3Zm5.65 0L7.4 13.05q-.275-.275-.275-.688t.275-.712q.3-.3.713-.3t.712.3l3.525 3.525l8.5-8.5q.3-.3.7-.287t.7.312q.275.3.288.7t-.288.7l-9.2 9.2q-.3.3-.7.3t-.7-.3Zm.7-4.95l-1.425-1.4l4.25-4.25q.275-.275.687-.275t.713.275q.3.3.3.713t-.3.712L12.35 12.35Z"/>`),
		g.Group(children),
	)
}