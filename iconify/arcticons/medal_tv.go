package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.061 27.188l5.856-3.474M24 27.758l12.939-7.388v15.064l6.561-4.102V16.476l-6.561-3.91L24 20.146l-12.939-7.58l-6.561 3.91v14.856l6.561 4.102V20.37L24 27.758z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.939 35.434L24 27.758l-12.939 7.676m25.878-8.246l-5.856-3.474"/>`),
		g.Group(children),
	)
}