package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallProtect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.69 0 15.25-7.77 15.25-16.94v-20c-4 0-15.25-2-15.25-2s-11.26 2-15.25 2v20C8.75 35.73 22.31 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.826 21.067c.587-.6 2.79-2.856 2.996-3.054a1.187 1.187 0 0 0 .219-1.144a23.072 23.072 0 0 1-.42-2.388a.764.764 0 0 0-.871-.67h-3.217a.568.568 0 0 0-.513.527a13.445 13.445 0 0 0 1.744 6.62h0l.014.025l.048.085l.002-.003a13.421 13.421 0 0 0 4.924 4.912l-.005.005l.165.092l.005.003h0a13.35 13.35 0 0 0 6.553 1.711a.568.568 0 0 0 .527-.513v-3.22a.765.765 0 0 0-.67-.872a23.034 23.034 0 0 1-2.386-.42a1.184 1.184 0 0 0-1.142.22c-.199.207-2.45 2.41-3.052 3"/>`),
		g.Group(children),
	)
}