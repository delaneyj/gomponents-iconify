package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricBikeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17q-2.125 0-3.563-1.438T0 12q0-2.125 1.463-3.563T5 7q1.925 0 3.238 1.15T9.9 11h.65l-1.8-5H7V4h5v2h-1.1l.35 1h4.8L14.6 3H12V1h2.6q.65 0 1.163.35t.737.95l1.7 4.65h.8q2.075 0 3.538 1.463T24 11.95q0 2.1-1.45 3.575T19 17q-1.8 0-3.163-1.125T14.1 13H9.9q-.35 1.725-1.7 2.863T5 17Zm0-2q1.025 0 1.763-.563T7.8 13H5v-2h2.8q-.3-.9-1.037-1.45T5 9q-1.275 0-2.138.863T2 12q0 1.25.863 2.125T5 15Zm7.7-4h1.4q.125-.575.338-1.075T15 9h-3.05l.75 2Zm6.3 4q1.275 0 2.138-.875T22 12q0-1.275-.863-2.138T19 9h-.1l1 2.65l-1.9.7l-.95-2.65q-.5.425-.775 1T16 12q0 1.25.863 2.125T19 15Zm-6 8l-6-3h4v-2l6 3h-4v2ZM4.9 12ZM19 12Z"/>`),
		g.Group(children),
	)
}