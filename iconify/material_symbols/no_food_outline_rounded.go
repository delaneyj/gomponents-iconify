package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoFoodOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.575L1.375 4.2q-.3-.3-.288-.7t.313-.7q.3-.3.713-.3t.712.3L21.2 21.175q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3Zm1.85-3.825L19.8 16.925L20.8 7h-9.55l-.1-.875q-.05-.45.238-.788T12.125 5H16V2q0-.425.288-.713T17 1q.425 0 .713.288T18 2v3h3.9q.45 0 .75.325t.25.775l-1.275 12.65Zm-6-5.975ZM2 19q-.425 0-.713-.288T1 18q0-.425.288-.713T2 17h13q.425 0 .713.288T16 18q0 .425-.288.713T15 19H2Zm0 4q-.425 0-.713-.288T1 22q0-.425.288-.713T2 21h13q.425 0 .713.288T16 22q0 .425-.288.713T15 23H2ZM9.05 9.025v2q-.125 0-.275-.013T8.5 11q-1.475 0-2.788.5T3.676 13h9.35l2 2H2.1q-.475 0-.775-.338T1.1 13.9q.475-2.475 2.65-3.688T8.5 9q.125 0 .275.013t.275.012ZM8.5 13Z"/>`),
		g.Group(children),
	)
}