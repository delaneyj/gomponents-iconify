package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M582 413v356q0 32-32 32H32q-32 0-32-32V413q0-33 32-33h389V186q0-51-37-74t-91-23q-53 0-92 25.5T162 191v125H65V188h1q9-87 74-140.5T293-6t151 63t71 151h1l2 172q4 0 14.5-.5t15 0t13 2t12 5t6.5 10t3 16.5z"/>`),
		g.Group(children),
	)
}