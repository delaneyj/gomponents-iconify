package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermostatAutoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.725 13q-.45 0-.662-.313t-.038-.737l2.7-7.2q.125-.3.425-.525T17.8 4q.325 0 .638.225t.437.525l2.75 7.225q.175.425-.038.725t-.687.3q-.175 0-.35-.113t-.225-.287l-.65-1.9h-3.7l-.65 1.9q-.05.175-.225.288t-.375.112Zm1.7-3.6h2.75L17.85 5.65h-.05L16.425 9.4ZM8 21q-2.075 0-3.538-1.463T3 16q0-1.2.525-2.238T5 12V6q0-1.25.875-2.125T8 3q1.25 0 2.125.875T11 6v6q.95.725 1.475 1.763T13 16q0 2.075-1.463 3.538T8 21Zm-3-5h6q0-.725-.313-1.35T9.8 13.6L9 13V6q0-.425-.288-.713T8 5q-.425 0-.713.288T7 6v7l-.8.6q-.575.425-.888 1.05T5 16Z"/>`),
		g.Group(children),
	)
}