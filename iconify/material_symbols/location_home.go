package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationHome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 15q1.25 0 2.125-.875T15 12q0-1.25-.875-2.125T12 9q-1.25 0-2.125.875T9 12q0 1.25.875 2.125T12 15Zm-6 4h12v-.275q0-.35-.15-.65t-.45-.475q-1.225-.775-2.587-1.188T12 16q-1.45 0-2.813.413T6.6 17.6q-.3.175-.45.475t-.15.65V19Zm-2 2V9l8-6l8 6v12H4Z"/>`),
		g.Group(children),
	)
}