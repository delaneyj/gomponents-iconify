package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 5H3l3.364-2.77A1 1 0 0 1 8 3.002v9.996a1 1 0 0 1-1.636.772L3 11H0V5h1.5ZM3 9.5H1.5v-3h2.038l.416-.342L6.5 4.061v7.878L3.954 9.842L3.538 9.5H3Zm9.457-7.267a.75.75 0 0 1 1.06-.026a7.999 7.999 0 0 1 0 11.586a.75.75 0 0 1-1.034-1.086a6.5 6.5 0 0 0 0-9.414a.75.75 0 0 1-.026-1.06Zm-2.292 2.034a.75.75 0 0 1 1.057-.09a5 5 0 0 1 .001 7.645a.75.75 0 1 1-.967-1.146a3.5 3.5 0 0 0 0-5.353a.75.75 0 0 1-.091-1.056Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}