package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PiggyBank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.5 14H20v-2h-2v-1h-2v1.05a2.5 2.5 0 0 0 .5 4.95h1a.5.5 0 0 1 0 1H14v2h2v1h2v-1.05a2.5 2.5 0 0 0-.5-4.95h-1a.5.5 0 0 1 0-1Z"/><path fill="currentColor" d="M29 13h-2.02A5.779 5.779 0 0 0 25 8.852V5a1 1 0 0 0-1.6-.8L19.667 7H15c-5.51 0-9.463 3.241-9.948 8H5a1 1 0 0 1-1-1v-2H2v2a3.003 3.003 0 0 0 3 3h.07A9.173 9.173 0 0 0 9 23.557V27a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-2h3v2a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-3.363A5.093 5.093 0 0 0 26.819 20H29a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1Zm-1 5h-2.876c-.305 2.753-.824 3.485-3.124 4.315V26h-2v-3h-7v3h-2v-3.622A7.013 7.013 0 0 1 7 16c0-4.835 4.018-7 8-7h5.333L23 7v2.776c2.418 1.86 1.913 3.186 2.018 5.224H28Z"/>`),
		g.Group(children),
	)
}