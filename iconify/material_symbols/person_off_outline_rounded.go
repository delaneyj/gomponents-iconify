package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.075 21.9L17.15 20H5q-.425 0-.713-.288T4 19v-1.8q0-.85.438-1.563T5.6 14.55q1.125-.575 2.288-.925t2.362-.525L2.075 4.925q-.3-.3-.287-.7t.312-.7q.3-.3.713-.3t.712.3L20.5 20.5q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3ZM6 18h9.15l-3-3H12q-1.4 0-2.775.338T6.5 16.35q-.225.125-.363.35T6 17.2v.8Zm12.4-3.45q.725.35 1.15 1.063T20 17.15l-3.35-3.35q.45.175.888.35t.862.4Zm-4.2-3.2l-1.475-1.475q.575-.225.925-.737T14 8q0-.825-.587-1.412T12 6q-.625 0-1.137.35t-.738.925L8.65 5.8q.575-.85 1.45-1.325T12 4q1.65 0 2.825 1.175T16 8q0 1.025-.475 1.9T14.2 11.35Zm.95 6.65H6h9.15Zm-3.725-9.425Z"/>`),
		g.Group(children),
	)
}