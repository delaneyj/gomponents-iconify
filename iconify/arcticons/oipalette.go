package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oipalette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.331 35.876s6.923-1.06 6.923-3.49c.186-2.1-2.378-1.07-2.458-2.993s9.705-7.444 5.771-13.25s-14.081-5.318-19.45-1.39s-8.963 11.79-4.808 17.419c3.008 3.932 6.115 4.56 14.022 3.704l-10.365-14.32l6.255-4.905c4.901-3.203 11.177 4.443 5.877 8.442l-6.095 4.803"/>`),
		g.Group(children),
	)
}