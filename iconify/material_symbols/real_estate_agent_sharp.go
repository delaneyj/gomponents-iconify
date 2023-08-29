package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RealEstateAgentSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 15v-2.4L9.35 9H7V6.5l7-5l7 5V15h-2ZM1 22V11h4v11H1Zm13 0l-7-1.975V11h1.975L17 14v2h-4l-1.75-.675l-.35.925L13 17h9v2l-8 3Zm.5-14h1V7h-1v1Zm-2 0h1V7h-1v1Zm2 2h1V9h-1v1Zm-2 0h1V9h-1v1Z"/>`),
		g.Group(children),
	)
}