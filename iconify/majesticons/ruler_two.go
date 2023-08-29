package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.879 3.707a3 3 0 0 1 4.242 0l3.172 3.172a3 3 0 0 1 0 4.242l-9.172 9.172a3 3 0 0 1-4.242 0L3.707 17.12a3 3 0 0 1 0-4.242L4.586 12l1.707 1.707a1 1 0 0 0 1.414-1.414L6 10.586L7.586 9l2.707 2.707a1 1 0 0 0 1.414-1.414L9 7.586L10.586 6l1.707 1.707a1 1 0 1 0 1.414-1.414L12 4.586l.879-.879z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}