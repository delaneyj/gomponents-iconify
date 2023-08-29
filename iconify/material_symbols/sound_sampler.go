package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundSampler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.225 18.325q-1.05-1.3-1.637-2.913T2 12q0-3.925 2.588-6.75T11 2.05v2q-3 .35-5 2.6T4 12q0 1.4.425 2.638T5.65 16.9l-1.425 1.425ZM12 22q-1.825 0-3.45-.6t-2.925-1.675l1.4-1.425q1.025.8 2.288 1.25T12 20q1.425 0 2.688-.45t2.287-1.25l1.4 1.425q-1.3 1.075-2.925 1.675T12 22Zm7.775-3.675L18.35 16.9q.8-1.025 1.225-2.263T20 12q0-3.1-2-5.35t-5-2.6v-2q3.825.375 6.413 3.2T22 12q0 1.8-.588 3.413t-1.637 2.912ZM9.5 16.5v-9l7 4.5l-7 4.5Z"/>`),
		g.Group(children),
	)
}