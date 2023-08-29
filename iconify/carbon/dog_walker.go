package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DogWalker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 20h2v10h-2zM4 23h2v7H4z"/><path fill="currentColor" d="M16 30h-2v-4a1.001 1.001 0 0 0-1-1h-2v5H9v-7h4a3.003 3.003 0 0 1 3 3zm-8.5-8A3.504 3.504 0 0 1 4 18.5V17H2v-2h4v3.5a1.5 1.5 0 0 0 3 0V15h4v2h-2v1.5A3.504 3.504 0 0 1 7.5 22zM27 10h-5.646a2.986 2.986 0 0 0-2.786 1.886l-1.442 3.605l-3.607 4.51l1.562 1.25l3.7-4.626l1.645-3.996a.995.995 0 0 1 .928-.629H27a1 1 0 0 1 1 1v7h-3v10h2v-8h1a2.002 2.002 0 0 0 2-2v-7a3.003 3.003 0 0 0-3-3zm-3-1a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}