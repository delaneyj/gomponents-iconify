package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 24h2v2H6zm2 2h2v2H8zm2 2h2v2h-2zm0-4h2v2h-2zm-4 4h2v2H6z"/><path fill="currentColor" d="M21 30a1 1 0 0 1-.894-1.447l2-4a1 1 0 1 1 1.788.894l-2 4A.998.998 0 0 1 21 30zm3.8-20.864a8.994 8.994 0 0 0-17.6 0A6.497 6.497 0 0 0 8.5 22h8.882l-1.276 2.553a1 1 0 1 0 1.788.894L19.619 22H23.5a6.497 6.497 0 0 0 1.3-12.864zM23.5 20h-15a4.497 4.497 0 0 1-.356-8.981l.816-.064l.099-.812a6.994 6.994 0 0 1 13.883 0l.099.812l.815.064A4.497 4.497 0 0 1 23.5 20z"/>`),
		g.Group(children),
	)
}