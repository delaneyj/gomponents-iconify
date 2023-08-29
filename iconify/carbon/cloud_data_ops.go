package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDataOps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 18h2v10h-2zm-4-4h2v14h-2zm-4 8h2v6h-2zm2.175-12l1.585-1.266a9.952 9.952 0 0 0-5.87-3.552a10.002 10.002 0 0 0-11.72 7.933A7.505 7.505 0 0 0 .054 21.41A7.684 7.684 0 0 0 7.77 28H16v-2H7.694a5.632 5.632 0 0 1-5.602-4.486a5.506 5.506 0 0 1 4.434-6.43l1.349-.245l.214-1.11a8.206 8.206 0 0 1 6.742-6.642a7.967 7.967 0 0 1 3.014.13A7.804 7.804 0 0 1 22.175 10z"/>`),
		g.Group(children),
	)
}