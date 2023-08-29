package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vehicle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 5a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v2h-2V5H2v7h6v2H2v4h6v2H5.414L3.5 21.914L2.086 20.5l.5-.5H2a2 2 0 0 1-2-2V5Zm11.323 3h10.354L24 13.807V21.5h-2V20H11v1.5H9v-7.693L11.323 8ZM11 18h11v-3.807L21.923 14H11.077l-.077.193V18Zm.877-6h9.246l-.8-2h-7.646l-.8 2ZM3 15h2.004v2.004H3V15Zm9 0h2.004v2.004H12V15Zm7 0h2.004v2.004H19V15Z"/>`),
		g.Group(children),
	)
}