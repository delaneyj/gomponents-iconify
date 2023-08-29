package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmokingNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm5 33a2 2 0 0 1 2-2h27a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-2Zm35-1a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4Zm-6-15a3 3 0 0 0-3 3v1.818a3 3 0 0 0 3 3h3a3 3 0 0 1 3 3V28a1 1 0 1 1-2 0v-.182a1 1 0 0 0-1-1h-3a5 5 0 0 1-5-5V20a5 5 0 0 1 5-5a1 1 0 1 1 0 2Zm4 2a1 1 0 1 0 0 2a4 4 0 0 1 4 4v3a1 1 0 1 0 2 0v-3a5.994 5.994 0 0 0-2.644-4.974A4.4 4.4 0 0 0 43 16.59V15a5 5 0 0 0-5-5a1 1 0 1 0 0 2a3 3 0 0 1 3 3v1.59A2.41 2.41 0 0 1 38.59 19H38Zm6 13a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0v-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}