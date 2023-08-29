package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpRealEstateAgent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 22h4V11H1v11zM14 1.5l-7 5V9h2l10 4v2h2V6.5l-7-5zm-.5 8.5h-1V9h1v1zm0-2h-1V7h1v1zm2 2h-1V9h1v1zm0-2h-1V7h1v1zM22 19l-8 3l-7-1.98V11h1.97L17 14v2h-4l-1.76-.68l-.33.94L13 17h9v2z"/>`),
		g.Group(children),
	)
}