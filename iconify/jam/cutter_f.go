package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CutterF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.684 17.995l-1.928 2.298L.588 19.5l4.5-5.362l4.596 3.857zm2.052-.89l-6.129-5.142l8.356-9.958a4 4 0 0 1 6.129 5.142l-8.356 9.959zm5.916-12.62a1 1 0 0 0-1.409.124l-.643.766a1 1 0 0 0 1.532 1.286l.643-.766a1 1 0 0 0-.123-1.41z"/>`),
		g.Group(children),
	)
}