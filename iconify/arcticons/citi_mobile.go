package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CitiMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.569 21.281v15.052m9.568-19.599v17.04a2.683 2.683 0 0 0 2.84 2.84h.852m-6.532-15.052h5.964m-18.253 12.22a5.497 5.497 0 0 1-4.828 2.84h0a5.697 5.697 0 0 1-5.68-5.68v-3.691a5.697 5.697 0 0 1 5.68-5.68h0a5.497 5.497 0 0 1 4.828 2.84m25.734-2.953V36.51"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.758 17.414c1.632-1.632 5.838-6.036 12.867-6.036S41.877 15.79 43.5 17.414"/>`),
		g.Group(children),
	)
}