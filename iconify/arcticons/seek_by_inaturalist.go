package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeekByInaturalist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 10.482c-1.96 45.487-30.732 17.697-39 26.716c2.213-44.666 32.18-16.956 39-26.716Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.216 28.784c6.405-9.858 14.99-9.467 23.82-7.915"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.397 16.324c4.931.338 6.996 2.515 8.639 4.545c-2.904.724-5.706 1.61-7.041 4.818"/>`),
		g.Group(children),
	)
}