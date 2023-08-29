package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RealEstateAgentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14V7.5L14 4L9 7.5V9H7V6.5l7-5l7 5V14h-2ZM14 4Zm.5 4h1V7h-1v1Zm-2 0h1V7h-1v1Zm2 2h1V9h-1v1Zm-2 0h1V9h-1v1ZM14 22.5l-7-1.95V22H1V11h7.95l6.2 2.3q.825.3 1.337 1.05T17 16h2q1.25 0 2.125.825T22 19v1l-8 2.5ZM3 20h2v-7H3v7Zm10.95.4l5.95-1.85q-.075-.275-.337-.413T19 18h-4.8q-.775 0-1.4-.1t-1.35-.35l-1.725-.6l.575-1.9l2 .65q.45.15 1.05.225T15 16q0-.275-.162-.525t-.388-.325L8.6 13H7v5.5l6.95 1.9ZM5 16.5Zm10-.5Zm-10 .5Zm2 0Z"/>`),
		g.Group(children),
	)
}