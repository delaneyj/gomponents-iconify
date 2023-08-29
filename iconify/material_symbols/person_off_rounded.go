package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.425 0-.713-.288T4 19v-1.8q0-.85.438-1.563T5.6 14.55q1.125-.575 2.288-.925t2.362-.525L2.075 4.925q-.3-.3-.287-.7t.312-.7q.3-.3.713-.3t.712.3L20.5 20.5q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3L17.15 20H5Zm15-2.85l-3.35-3.35q.45.175.888.35t.862.4q.725.35 1.15 1.063T20 17.15Zm-5.8-5.8L8.65 5.8q.575-.85 1.45-1.325T12 4q1.65 0 2.825 1.175T16 8q0 1.025-.475 1.9T14.2 11.35Z"/>`),
		g.Group(children),
	)
}