package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webtube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.916 9.733C-9.22 38.03 38.32 55.347 45.493 24.543"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.907 2.5c-31.58 17 6.357 52.346 21.586 22.043"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.478 33.05c.025.016 8.145-4.148 11.479-5.901c1.214-.638 1.15-1.684-.05-2.29a3784.438 3784.438 0 0 0-11.499-5.762a1.277 1.277 0 0 0-1.699 1.354c-.01 1.216 0 10.895 0 10.895"/>`),
		g.Group(children),
	)
}