package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fiftheditioncharactersheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 4.5l16.89 9.75v19.5L24 43.5L7.11 33.75v-19.5L24 4.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 13.17L13.74 30.93h20.52L24 13.17z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.26 30.93l6.63-16.68L24 13.17l10.26 17.76zM24 13.17L7.11 14.25l6.63 16.68L24 13.17zM13.74 30.93L24 43.5L7.11 33.75l6.63-2.82zm20.52 0L24 43.5l16.89-9.75l-6.63-2.82zM24 4.5v8.67"/>`),
		g.Group(children),
	)
}