package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveDoneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.125 21.1l-5.9-5.9l-2.15 2.15q-.3.3-.7.3t-.7-.3l-4.25-4.25q-.3-.3-.287-.7t.287-.7q.3-.3.713-.312t.712.287l3.525 3.525l1.4-1.4l-9.65-9.65q-.3-.3-.288-.7t.288-.7q.3-.3.713-.313t.712.288L22.525 19.7q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275Zm-15.1-3.75l-4.25-4.25q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l3.55 3.55l1.4 1.4l-.7.7q-.3.3-.7.3t-.7-.3Zm12-4.95l-1.4-1.4l4.225-4.225q.275-.275.675-.287t.7.262q.3.275.313.7t-.288.725L18.025 12.4Zm-2.85-2.85l-1.4-1.4l1.45-1.45q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-1.45 1.45Z"/>`),
		g.Group(children),
	)
}