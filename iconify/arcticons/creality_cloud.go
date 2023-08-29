package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrealityCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.833 6.142l9.291 6.538l8.947-6.538l9.429 6.47l-18.376 13.076L5.5 12.612l9.333-6.47Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.124 25.688v16.17L5.5 29.327V12.61m37 .001v16.716l-18.376 12.53"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.312 35.592V19.15l-9.188-6.47l-9.312 6.47v16.442"/>`),
		g.Group(children),
	)
}