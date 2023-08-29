package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devrant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.92 4.5a26.15 26.15 0 0 0-4 .27A15.07 15.07 0 0 0 12 8.63c-3 2.86-4.48 6.55-4.7 11.62c-.25 5.49 1.22 10 4.29 13.16a14.13 14.13 0 0 0 5.27 3.43l.57.2l1.69 2.59c.93 1.42 1.8 2.79 1.94 3c.54.94.8 1.06 1.06.47c.13-.27 1.18-4.81 1.18-5.08a3.67 3.67 0 0 1 .87-.05c6.63 0 11.62-2.56 14.42-7.34a17.79 17.79 0 0 0 2.18-8.6C40.89 17.5 40 14 38 11a14.72 14.72 0 0 0-10.1-6.24a25.33 25.33 0 0 0-4-.23Zm7.44 7.56l-7.18 18.05"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.62 26.16a2 2 0 0 1 2 2h0a2 2 0 1 1-2-2Zm0-9.47a2 2 0 0 1 2 2h0a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2h0a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}