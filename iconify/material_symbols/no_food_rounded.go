package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoFoodRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.625 18.75L11.45 8.6l-.3-2.475q-.05-.45.238-.788T12.125 5H16V2q0-.425.288-.713T17 1q.425 0 .713.288T18 2v3h3.9q.45 0 .75.325t.25.775l-1.275 12.65ZM2 19q-.425 0-.713-.288T1 18q0-.425.288-.713T2 17h13q.425 0 .713.288T16 18q0 .425-.288.713T15 19H2Zm0 4q-.425 0-.713-.288T1 22q0-.425.288-.713T2 21h13q.425 0 .713.288T16 22q0 .425-.288.713T15 23H2Zm17.775-.425L12.175 15H2.1q-.475 0-.775-.338T1.1 13.9q.475-2.475 2.65-3.688T8.5 9q.125 0 .263.013t.287.012l-.025 2.825l-7.65-7.65q-.3-.3-.288-.713t.313-.712q.3-.3.713-.3t.712.3l18.375 18.4q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3Z"/>`),
		g.Group(children),
	)
}