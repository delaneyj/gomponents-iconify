package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrontLoaderSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.975 19.775q-1.25 0-2.113-.875T1 16.775v-8h7v-5h5l3 3l-.025 4h1v-2l2.525-2.5l4.475 8.5h-7v-2h-1v1.75q.5.425.763 1.012T17 16.776q0 1.25-.875 2.125T14 19.775q-.95 0-1.725-.55t-1.125-1.45H6.825q-.35.9-1.125 1.45t-1.725.55Zm0-2q.425 0 .725-.287t.3-.713q0-.425-.288-.712T4 15.775q-.425 0-.713.288T3 16.774q0 .425.275.713t.7.287Zm10.025 0q.425 0 .713-.287t.287-.713q0-.425-.288-.712T14 15.775q-.425 0-.713.288t-.287.712q0 .425.288.713t.712.287Zm6.675-5L19 9.6v3.175h1.675Zm-10.675-2h4V7.6l-1.85-1.825H10v5Z"/>`),
		g.Group(children),
	)
}