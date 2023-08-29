package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doubtnut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.06 24.428a4.369 4.369 0 1 1-4.374 4.373a4.372 4.372 0 0 1 4.373-4.373Zm6.56-19.145l18.88 7.465m-5.998-2.125l-2.125 5.198a13.439 13.439 0 0 0-8.308-3.547l2.125-5.2m11.329 4.576l1.66 5.928"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.193 16.26h3.381l-.061 12.998a13.506 13.506 0 1 1-10.547-13.226v3.162a10.02 10.02 0 1 0 7.289 10.028Z"/>`),
		g.Group(children),
	)
}