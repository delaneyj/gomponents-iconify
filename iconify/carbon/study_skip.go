package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudySkip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M22 9h5.2A11 11 0 0 0 7 15H5a13 13 0 0 1 23-8.3V3h2v8h-8z" fill="currentColor"/><path d="M16 12a3 3 0 1 1-3 3a2.95 2.95 0 0 1 3-3m0-2a5 5 0 1 0 5 5a5 5 0 0 0-5-5z" fill="currentColor"/><path d="M23 30h-2v-3a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3v3H9v-3a5 5 0 0 1 5-5h4a5 5 0 0 1 5 5z" fill="currentColor"/>`),
		g.Group(children),
	)
}