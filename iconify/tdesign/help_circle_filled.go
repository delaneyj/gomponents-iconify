package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11Zm-.174-11.11c.432-.53.974-.974 1.41-1.318a2 2 0 1 0-3.123-2.24l-.332.944l-1.886-.666l.333-.943a4.001 4.001 0 1 1 6.246 4.476c-.431.34-.817.666-1.096 1.009c-.274.338-.378.61-.378.848v1.25h-2V14c0-.867.39-1.573.826-2.11ZM11 18.254V16.25h2.004v2.004H11Z"/>`),
		g.Group(children),
	)
}