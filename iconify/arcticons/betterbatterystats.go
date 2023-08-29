package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Betterbatterystats(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.5 4.5h6.67a1 1 0 0 1 1 1v3.004h4.635a2 2 0 0 1 2 2V41.5a2 2 0 0 1-2 2h-17.61a2 2 0 0 1-2-2V10.504a2 2 0 0 1 2-2H19.5V5.5a1 1 0 0 1 1-1ZM13.194 16h21.611M19.5 8.504h8.669M24 36.481v-13m-6.5 6.5h13"/>`),
		g.Group(children),
	)
}