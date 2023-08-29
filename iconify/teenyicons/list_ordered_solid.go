package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListOrderedSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 2h2v3h1v1H0V5h1V3H0V2Zm15 2H5V3h10v1Zm0 4H5V7h10v1ZM0 11.5A1.5 1.5 0 0 1 1.5 10h.293a1.207 1.207 0 0 1 .853 2.06l-.939.94H3v1H.5a.5.5 0 0 1-.354-.854l1.793-1.792A.207.207 0 0 0 1.793 11H1.5a.5.5 0 0 0-.5.5H0Zm15 .5H5v-1h10v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}