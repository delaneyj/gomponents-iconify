package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Style(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.975 19.8l-.85-.35q-.775-.325-1.038-1.125t.088-1.575l1.8-3.9v6.95Zm4 2.2q-.825 0-1.413-.588T5.976 20v-6l2.65 7.35q.075.175.15.338t.2.312h-1Zm5.15-.1q-.8.3-1.55-.075t-1.05-1.175l-4.45-12.2q-.3-.8.05-1.562t1.15-1.038l7.55-2.75q.8-.3 1.55.075t1.05 1.175l4.45 12.2q.3.8-.05 1.563t-1.15 1.037l-7.55 2.75ZM10.975 10q.425 0 .713-.288T11.975 9q0-.425-.287-.713T10.975 8q-.425 0-.712.288T9.975 9q0 .425.288.713t.712.287Z"/>`),
		g.Group(children),
	)
}