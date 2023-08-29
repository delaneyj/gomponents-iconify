package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Locked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 6.5V10H2a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2h-1.5V6.5a6.5 6.5 0 1 0-13 0zM6 10V6.5a4 4 0 0 1 8 0V10zm2 5.5a2 2 0 1 1 3.092 1.676l-.008.005s.195 1.18.415 2.57v.001a.749.749 0 0 1-.749.749H9.248a.749.749 0 0 1-.749-.749v-.001l.415-2.57a2.002 2.002 0 0 1-.916-1.68z"/>`),
		g.Group(children),
	)
}