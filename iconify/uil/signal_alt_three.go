package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalAltThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 9h-4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V10a1 1 0 0 0-1-1Zm-1 12h-2V11h2Zm9-20h-4a1 1 0 0 0-1 1v20a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Zm-1 20h-2V3h2ZM6 15H2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1Zm-1 6H3v-4h2Z"/>`),
		g.Group(children),
	)
}