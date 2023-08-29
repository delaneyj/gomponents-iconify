package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deploy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24.01 20.027v2h-24v-2h4v-1a2.006 2.006 0 0 1-2-2v-10a2.006 2.006 0 0 1 2-2h1.996v2H4.01v10h16v-10h-2.004v-2h2.004a2.006 2.006 0 0 1 2 2l-.01 10a1.997 1.997 0 0 1-1.99 2v1Zm-9-6.012l-3-3l-3 3h2v2.01h2v-2.01Zm.995-7.991a4 4 0 1 1-4-4a4.001 4.001 0 0 1 4 4Zm-4.4 2.96v-.56a.802.802 0 0 1-.8-.8v-.4L9.06 5.479a2.958 2.958 0 0 0 2.545 3.505Zm2.658-1.007a2.977 2.977 0 0 0-1.068-4.704a.797.797 0 0 1-.79.75h-.8v.8a.401.401 0 0 1-.4.4h-.8v.8h2.4a.401.401 0 0 1 .4.4v1.2h.4a.787.787 0 0 1 .658.354Z"/>`),
		g.Group(children),
	)
}