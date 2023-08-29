package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainHeavy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 30a1 1 0 0 1-.894-1.447l2-4a1 1 0 1 1 1.788.894l-2 4A.998.998 0 0 1 17 30zm-9 0a1 1 0 0 1-.894-1.447l2-4a1 1 0 1 1 1.788.894l-2 4A.998.998 0 0 1 8 30z"/><path fill="currentColor" d="M30 15.5a6.532 6.532 0 0 0-5.2-6.364a8.994 8.994 0 0 0-17.6 0a6.49 6.49 0 0 0-1.497 12.222l-1.597 3.195a1 1 0 1 0 1.788.894L7.617 22h6.765l-1.276 2.553a1 1 0 1 0 1.788.894L16.619 22h6.764l-1.276 2.553a1 1 0 1 0 1.788.894l1.945-3.89A6.506 6.506 0 0 0 30 15.5ZM23.5 20h-15a4.497 4.497 0 0 1-.356-8.981l.816-.064l.099-.812a6.994 6.994 0 0 1 13.883 0l.099.812l.815.064A4.497 4.497 0 0 1 23.5 20Z"/>`),
		g.Group(children),
	)
}