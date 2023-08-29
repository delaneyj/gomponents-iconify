package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.825 20h10.35L12 14.825L6.825 20Zm13.25 2.9l-.9-.9H4.425q-.675 0-.937-.613T3.7 20.3l6.9-6.9l-7.5-7.475q-.3-.275-.287-.688t.287-.712q.3-.3.713-.3t.712.3L21.5 21.5q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288ZM22 19.175l-2-2V6.825L14.825 12L13.4 10.6l6.9-6.9q.475-.475 1.088-.213t.612.938v14.75Zm-4.575-4.6ZM14.6 17.4Z"/>`),
		g.Group(children),
	)
}