package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ukrposhta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.304 24.44c10.34-2.145 7.439-14.555-.177-14.555c-7.854 0-10.674 11.23-1.822 14.084l13.23 3.11L24.95 43.5s-9.611-13.066-12.436-17.67C6.804 16.52 13.748 4.545 24.961 4.5c12.566 0 17.54 14 8.383 21.83"/>`),
		g.Group(children),
	)
}