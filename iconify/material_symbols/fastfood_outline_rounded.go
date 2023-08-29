package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastfoodOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.15 15q-.475 0-.775-.338T1.15 13.9q.475-2.475 2.65-3.688T8.55 9q2.575 0 4.75 1.213t2.65 3.687q.075.425-.225.763T14.95 15H2.15Zm15.9 8v-2h1.4l1.4-14H11.3l-.1-.875q-.05-.45.238-.788T12.175 5h3.875V2q0-.425.288-.713T17.05 1q.425 0 .713.288T18.05 2v3h3.9q.45 0 .75.325t.25.775L21.4 21.55q-.075.6-.537 1.025T19.75 23h-1.7Zm0-2h1.4h-1.4Zm-4.725-8q-.725-1-2.037-1.5T8.5 11q-1.475 0-2.788.5T3.676 13h9.65ZM8.5 13ZM2 19q-.425 0-.713-.288T1 18q0-.425.288-.713T2 17h13q.425 0 .713.288T16 18q0 .425-.288.713T15 19H2Zm0 4q-.425 0-.713-.288T1 22q0-.425.288-.713T2 21h13q.425 0 .713.288T16 22q0 .425-.288.713T15 23H2Z"/>`),
		g.Group(children),
	)
}