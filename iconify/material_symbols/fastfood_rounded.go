package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastfoodRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.15 15q-.475 0-.775-.338T1.15 13.9q.475-2.475 2.65-3.688T8.55 9q2.575 0 4.75 1.213t2.65 3.687q.075.425-.225.763T14.95 15H2.15Zm15.9 8v-8q0-2.875-1.763-4.888T11.326 7.3L11.2 6.125q-.05-.45.238-.787T12.175 5h3.875V2q0-.425.288-.713T17.05 1q.425 0 .713.288T18.05 2v3h3.9q.45 0 .75.325t.25.775L21.4 21.55q-.075.6-.537 1.025T19.75 23h-1.7ZM2 19q-.425 0-.713-.288T1 18q0-.425.288-.713T2 17h13q.425 0 .713.288T16 18q0 .425-.288.713T15 19H2Zm0 4q-.425 0-.713-.288T1 22q0-.425.288-.713T2 21h13q.425 0 .713.288T16 22q0 .425-.288.713T15 23H2Z"/>`),
		g.Group(children),
	)
}