package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.8 12.136a8.87 8.87 0 0 0-.979-2.543L30 3.414L28.587 2L2 28.587L3.414 30l5-5H23.5a6.497 6.497 0 0 0 1.3-12.864zM23.5 23H10.414l11.928-11.928a6.905 6.905 0 0 1 .6 2.071l.099.812l.815.064A4.497 4.497 0 0 1 23.5 23zm-19.204.449l1.432-1.431a4.477 4.477 0 0 1 2.416-7.999l.816-.064l.099-.812a6.987 6.987 0 0 1 10.63-5.086l1.443-1.443A8.986 8.986 0 0 0 7.2 12.136A6.49 6.49 0 0 0 4.296 23.45z"/>`),
		g.Group(children),
	)
}