package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BetterUk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><ellipse cx="24.038" cy="14.088" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.064" ry="4.588"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.554 21.673h18.892M20.166 38.5V21.673M27.834 38.5V21.673"/>`),
		g.Group(children),
	)
}