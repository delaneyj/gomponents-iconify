package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.956 28.188V19.82h2.74a2.81 2.81 0 0 1 0 5.62h-2.74m2.74.001l2.739 2.745M16.852 19.82h5.544m-2.772 8.368V19.82m-9.958 7.451a2.34 2.34 0 0 0 2.051.917h1.24a2.09 2.09 0 0 0 2.087-2.092h0a2.09 2.09 0 0 0-2.088-2.092h-1.368A2.09 2.09 0 0 1 9.5 21.912h0a2.09 2.09 0 0 1 2.088-2.092h1.239a2.34 2.34 0 0 1 2.051.917m11.674-.925l1.04 3.199h3.364l-2.722 1.978l1.04 3.199l-2.722-1.977l-2.721 1.977l1.039-3.199l-2.722-1.978h3.365l1.039-3.199z"/>`),
		g.Group(children),
	)
}