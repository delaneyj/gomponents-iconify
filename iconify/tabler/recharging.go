package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Recharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.038 4.5a9 9 0 0 0-2.495 2.47m-1.357 3.239a9 9 0 0 0 0 3.508M4.5 16.962a9 9 0 0 0 2.47 2.495m3.239 1.357a9 9 0 0 0 3.5 0m3.253-1.314a9 9 0 0 0 2.495-2.47m1.357-3.239a9 9 0 0 0 0-3.508M19.5 7.038a9 9 0 0 0-2.47-2.495m-3.239-1.357a9 9 0 0 0-3.508-.02M12 8l-2 4h4l-2 4"/><path d="M12 21a9 9 0 0 0 0-18"/></g>`),
		g.Group(children),
	)
}