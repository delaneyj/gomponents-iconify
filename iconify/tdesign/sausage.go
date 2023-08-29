package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sausage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.998 2.003v1.996h2.004v2h-1.257a4.225 4.225 0 0 1-1.04 4.268l-9.44 9.44a4.225 4.225 0 0 1-4.267 1.04V22h-2v-2H2.002v-2h1.249a4.225 4.225 0 0 1 1.04-4.269l9.44-9.44a4.225 4.225 0 0 1 4.267-1.04v-1.25h2Zm-1.707 3.704a2.225 2.225 0 0 0-3.146 0l-9.44 9.44a2.225 2.225 0 1 0 3.146 3.146l9.44-9.44a2.225 2.225 0 0 0 0-3.146Z"/>`),
		g.Group(children),
	)
}