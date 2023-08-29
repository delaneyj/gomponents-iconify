package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1920 1024q0 159-112.5 271.5T1536 1408H448q-185 0-316.5-131.5T0 960q0-132 71-241.5T258 555q-2-28-2-43q0-212 150-362T768 0q158 0 286.5 88T1242 318q70-62 166-62q106 0 181 75t75 181q0 75-41 138q129 30 213 134.5t84 239.5z"/>`),
		g.Group(children),
	)
}