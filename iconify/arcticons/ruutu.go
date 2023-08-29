package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruutu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.484 15.826L5.005 28.61c-.977.98-.614 4.145 1.777 4.205c3.17.08 3.003-.042 4.854-.042m9.157-14.996c0-1.636-2.555-3.699-4.291-1.96m-2.084 9.786c0-1.4 1.133-2.536 2.53-2.536m-2.53 0v6.72m20.556-8.812v8.812m-1.328-6.72h2.657m-17.17 0v4.184c0 1.4 1.129 2.536 2.52 2.536s2.52-1.136 2.52-2.536v-4.184m.001 4.184v2.536m2.398-6.72v4.184c0 1.4 1.129 2.536 2.52 2.536s2.52-1.136 2.52-2.536v-4.184m.001 4.184v2.536m6.847-6.72v4.184c0 1.4 1.128 2.536 2.52 2.536s2.52-1.136 2.52-2.536v-4.184m0 4.184v2.536"/>`),
		g.Group(children),
	)
}