package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorkspacePremiumOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.925 12.75L12 11.925l1.05.825q.275.225.6.013t.2-.563l-.425-1.35l1.2-.95q.275-.225.163-.563T14.325 9H12.9l-.425-1.325q-.125-.35-.475-.35t-.475.35L11.1 9H9.675q-.35 0-.475.35t.15.55l1.2.95l-.425 1.35q-.125.35.188.563t.612-.013Zm-3.6 9.8q-.5.175-.913-.125t-.412-.8v-6.35q-.95-1.05-1.475-2.4T4 10q0-3.35 2.325-5.675T12 2q3.35 0 5.675 2.325T20 10q0 1.525-.525 2.875T18 15.275v6.35q0 .5-.413.8t-.912.125L12 21l-4.675 1.55ZM12 16q2.5 0 4.25-1.75T18 10q0-2.5-1.75-4.25T12 4Q9.5 4 7.75 5.75T6 10q0 2.5 1.75 4.25T12 16Zm-4 4.025L12 19l4 1.025v-3.1q-.875.5-1.888.788T12 18q-1.1 0-2.113-.288T8 16.925v3.1Zm4-1.55Z"/>`),
		g.Group(children),
	)
}