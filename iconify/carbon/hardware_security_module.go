package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardwareSecurityModule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="21.5" cy="7.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M14.414 20H9v-5.414l6.03-6.03A5.352 5.352 0 0 1 15 8a6 6 0 1 1 6 6a5.358 5.358 0 0 1-.556-.03zM11 18h2.586l6.17-6.17l.518.095A3.935 3.935 0 0 0 21 12a4.05 4.05 0 1 0-3.925-3.273l.095.517l-6.17 6.17zm17 2h-9v2h9v6H4v-6h2v-2H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2z"/><circle cx="7" cy="25" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}