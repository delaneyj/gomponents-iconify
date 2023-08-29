package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdidasTraining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.227 12.859l-7.566 4.332l10.523 17.95H43.5L30.227 12.859zm-9.138 7.765l-7.565 4.333l5.97 10.184h10.243l-8.648-14.517zm-9.024 7.259L4.5 32.216l1.715 2.925h10.174l-4.324-7.258z"/>`),
		g.Group(children),
	)
}