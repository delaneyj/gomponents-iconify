package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AcuteOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 20q-3.35 0-5.675-2.325T7 12q0-3.325 2.325-5.663T15 4q3.325 0 5.663 2.337T23 12q0 3.35-2.337 5.675T15 20Zm0-2q2.5 0 4.25-1.75T21 12q0-2.5-1.75-4.25T15 6q-2.5 0-4.25 1.75T9 12q0 2.5 1.75 4.25T15 18Zm2.275-2.275L18.7 14.3L16 11.6V8h-2v4.425l3.275 3.3ZM2 9V7h4v2H2Zm-1 4v-2h5v2H1Zm1 4v-2h4v2H2Zm13-5Z"/>`),
		g.Group(children),
	)
}