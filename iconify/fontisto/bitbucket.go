package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitbucket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.865 0H.854A.854.854 0 0 0 0 .843v.011c0 .048.004.095.012.141L.011.99l3.63 22.041c.096.55.567.964 1.136.97h17.424a.859.859 0 0 0 .847-.714l.001-.005l3.638-22.281a.856.856 0 0 0-.701-.98L25.981.02a.747.747 0 0 0-.129-.011h-.02h.001zm15.287 15.926H10.59L9.088 8.068h8.411z"/>`),
		g.Group(children),
	)
}