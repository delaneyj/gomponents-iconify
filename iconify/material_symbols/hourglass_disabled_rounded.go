package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassDisabledRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.775 22.575l-.6-.575H5q-.425 0-.713-.288T4 21q0-.425.288-.713T5 20h1v-3q0-1.525.713-2.863T8.7 12q-.8-.5-1.362-1.2t-.913-1.55L1.375 4.2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l18.4 18.4q.3.3.3.7t-.3.7q-.3.3-.712.3t-.713-.3ZM15.05 12.15l-1.5-1.475q1.125-.475 1.788-1.475T16 7V4H8v1.125L6.875 4l-2-2H19q.425 0 .713.288T20 3q0 .425-.288.713T19 4h-1v3q0 1.6-.775 2.975T15.05 12.15ZM8 20h8v-1.175L10.475 13.3q-1.125.475-1.8 1.475T8 17v3Zm10 0Z"/>`),
		g.Group(children),
	)
}