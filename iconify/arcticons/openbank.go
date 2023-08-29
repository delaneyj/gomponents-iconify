package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.686 18.52a8.15 8.15 0 1 0 1.911 7.337a8.652 8.652 0 0 0 .201-1.605a2.323 2.323 0 0 1 1.432-2.138l.014-.006a2.258 2.258 0 0 1 .825-.156H43.5m-6.994 6.702a3.497 3.497 0 1 0 6.994 0v-6.702"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.512 26.552v2.102a3.497 3.497 0 1 0 6.994 0v-2.102"/>`),
		g.Group(children),
	)
}