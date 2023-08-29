package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.1.7a.5.5 0 0 1 .4-.2h9a.5.5 0 0 1 .4.2l2.976 3.974c.149.185.156.45.01.644L8.4 15.3a.5.5 0 0 1-.8 0L.1 5.3a.5.5 0 0 1 0-.6l3-4zm11.386 3.785l-1.806-2.41l-.776 2.413l2.582-.003zm-3.633.004l.961-2.989H4.186l.963 2.995l5.704-.006zM5.47 5.495L8 13.366l2.532-7.876l-5.062.005zm-1.371-.999l-.78-2.422l-1.818 2.425l2.598-.003zM1.499 5.5l5.113 6.817l-2.192-6.82L1.5 5.5zm7.889 6.817l5.123-6.83l-2.928.002l-2.195 6.828z"/>`),
		g.Group(children),
	)
}