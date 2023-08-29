package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.086 16.288l2.633-1.52l.197-1.923l2.381-1.375l.341 2l-.445 5.784l6.526 11.304v4.71l-4.079-2.355l-6.526-11.303l-5.23-2.507l-1.563-1.295l2.38-1.375l1.765.79l2.633-1.52m6.541 17.21l4.08-2.355M20.114 21.61l-2.591-4.489m6.67 2.134l-2.468-4.275m.871 10.908l3.902-2.252"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.267 15.67v16.66a3.84 3.84 0 0 1-1.92 3.325l-14.427 8.33a3.84 3.84 0 0 1-3.84 0l-14.428-8.33a3.84 3.84 0 0 1-1.92-3.325V15.67a3.84 3.84 0 0 1 1.92-3.325l14.428-8.33a3.84 3.84 0 0 1 3.84 0l14.428 8.33a3.84 3.84 0 0 1 1.92 3.325Z"/>`),
		g.Group(children),
	)
}