package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Joyn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><circle cx="18.257" cy="21.386" r="3.637" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 17.75h2.967v9.762a2.764 2.764 0 0 1-2.745 2.782q-.111.001-.222-.007M24.334 17.75v5.152c0 1.451 2.107 3.94 5.456 1.116"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.79 17.75v9.857c0 3.637-5.456 3.063-5.456.479m8.328-10.337v7.557m0-4.686c1.531-2.87 5.838-4.307 5.838-1.052v5.738"/>`),
		g.Group(children),
	)
}