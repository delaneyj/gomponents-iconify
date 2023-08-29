package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiningSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19h1.5v-6.75q.65-.2 1.075-.713T11 10.3V6.5h-1V9h-.75V6.5h-1V9H7.5V6.5h-1v3.8q0 .725.425 1.238T8 12.25V19Zm6 0h1.5v-6.35q.825-.4 1.288-1.275t.462-2.05q0-1.425-.713-2.375T14.75 6q-1.075 0-1.788.95t-.712 2.375q0 1.175.463 2.05T14 12.65V19ZM2 22V2h20v20H2Z"/>`),
		g.Group(children),
	)
}