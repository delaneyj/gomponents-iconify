package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeSpeakerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.575 21Q7.75 21 6.55 19.637t-.95-3.162l1.375-10.3q.05-.275.2-.487t.425-.313l7.9-3.15q.45-.2.875.063t.475.762l1.6 13.475q.225 1.8-.962 3.138T14.5 21H9.575Zm-1.35-9H15.9L15 4.575l-6.1 2.45L8.225 12Zm1.35 7H14.5q.9 0 1.488-.675t.487-1.575L16.15 14h-8.2l-.35 2.725q-.125.9.475 1.588t1.5.687Zm2.475-5v-2v2Zm.025-2Zm-.05 2Z"/>`),
		g.Group(children),
	)
}