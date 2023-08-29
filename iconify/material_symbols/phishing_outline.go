package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhishingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21q-2.5 0-4.25-1.75T5 15V9l5 5l-1.4 1.425l-1.6-1.6V15q0 1.65 1.175 2.825T11 19q1.65 0 2.825-1.175T15 15v-3.175q-.9-.35-1.45-1.112T13 9q0-.95.55-1.713T15 6.175V2h2v4.175q.9.35 1.45 1.113T19 9q0 .95-.55 1.725t-1.45 1.1V15q0 2.5-1.75 4.25T11 21Zm5-11q.425 0 .713-.288T17 9q0-.425-.288-.713T16 8q-.425 0-.713.288T15 9q0 .425.288.713T16 10Zm0-1Z"/>`),
		g.Group(children),
	)
}