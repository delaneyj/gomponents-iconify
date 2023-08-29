package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BehanceAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 6h5a.5.5 0 0 0 0-1h-5a.5.5 0 0 0 0 1zm7.026 9.003a.494.494 0 0 0-.553.441A3.999 3.999 0 0 1 14 15v-2h8.47a.5.5 0 0 0 .497-.553A5.004 5.004 0 0 0 18 8a5.006 5.006 0 0 0-5 5v2a4.999 4.999 0 0 0 9.967.556a.5.5 0 0 0-.44-.553zM18 9a4.015 4.015 0 0 1 3.874 3h-7.747A4.006 4.006 0 0 1 18 9zm-8.604 2.434A3.495 3.495 0 0 0 7.5 5h-6a.5.5 0 0 0-.5.5v14a.5.5 0 0 0 .5.5h6a4.491 4.491 0 0 0 1.896-8.566zM2 6h5.5a2.5 2.5 0 1 1 0 5H2V6zm5.5 13H2v-7h5.5a3.5 3.5 0 0 1 0 7z"/>`),
		g.Group(children),
	)
}