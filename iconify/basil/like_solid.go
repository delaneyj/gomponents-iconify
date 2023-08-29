package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LikeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.337 3.75a1.4 1.4 0 0 0-1.149.6l-4.232 6.077a5.75 5.75 0 0 0-.997 3.914l.326 2.961a2.736 2.736 0 0 0 2.397 2.419l2.117.252c1.616.192 3.255.025 4.8-.489a3.846 3.846 0 0 0 2.365-2.24l1.709-4.343A2.818 2.818 0 0 0 15.43 9.12l-3.898.878l1.16-4.499a1.4 1.4 0 0 0-1.356-1.749Z"/>`),
		g.Group(children),
	)
}