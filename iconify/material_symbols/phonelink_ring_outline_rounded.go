package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkRingOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.525 12q0 .55-.137 1.063t-.413.987q-.2.35-.625.363t-.725-.288q-.3-.3-.313-.713t.113-.837q.05-.125.075-.275t.025-.3q0-.15-.025-.3t-.075-.275q-.125-.425-.113-.838t.313-.712q.3-.3.713-.3t.612.35q.275.475.425.987t.15 1.088Zm3.5 0q0 1.25-.4 2.413t-1.15 2.137q-.25.325-.663.338T20.1 16.6q-.275-.275-.288-.7t.238-.775q.475-.675.725-1.475t.25-1.65q0-.85-.25-1.65t-.725-1.475q-.25-.35-.238-.775t.288-.7q.3-.3.713-.287t.662.337q.75.975 1.15 2.138t.4 2.412ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v4h-2V6H7v12h10v-1h2v4q0 .825-.588 1.413T17 23H7Zm0-3v1h10v-1H7ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}