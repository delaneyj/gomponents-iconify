package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncognitoTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 7.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0 0l.211-.106a4 4 0 0 1 3.578 0L14 7.5m0 0a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0Zm-2 6.303c5-3 5 3.5 9 1.767c-1 4.233-6 4.233-9 1.233c-3 3-8 3-9-1.233c4 1.733 4-4.767 9-1.767Z"/>`),
		g.Group(children),
	)
}