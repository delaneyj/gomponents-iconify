package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 22a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2H6V6.5a4 4 0 0 1 8 0v.282a1.25 1.25 0 1 0 2.5 0v-.033v.002v-.25a6.5 6.5 0 1 0-13 0v3.5H2a2 2 0 0 0-2 2v10zm8-6.5a2 2 0 1 1 3.092 1.676l-.008.005s.195 1.18.415 2.57v.001a.749.749 0 0 1-.749.749H9.248a.749.749 0 0 1-.749-.749v-.001l.415-2.57a2.002 2.002 0 0 1-.916-1.68z"/>`),
		g.Group(children),
	)
}