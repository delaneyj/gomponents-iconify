package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhishingOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21q-2.5 0-4.25-1.75T5 15v-4.8q0-.35.3-.475t.55.125L9.3 13.3q.275.275.275.688T9.3 14.7q-.3.3-.713.3t-.712-.3L7 13.825V15q0 1.65 1.175 2.825T11 19q1.65 0 2.825-1.175T15 15v-3.175q-.9-.35-1.45-1.112T13 9q0-.95.55-1.713T15 6.175V3q0-.425.288-.713T16 2q.425 0 .713.288T17 3v3.175q.9.35 1.45 1.113T19 9q0 .95-.55 1.725t-1.45 1.1V15q0 2.5-1.75 4.25T11 21Zm5-11q.425 0 .713-.288T17 9q0-.425-.288-.713T16 8q-.425 0-.713.288T15 9q0 .425.288.713T16 10Zm0-1Z"/>`),
		g.Group(children),
	)
}