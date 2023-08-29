package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitRepositoryCommits(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5.5A4.5 4.5 0 0 1 7.5 1H21v22h-5v-2h3v-5h-2.5v-2H19V3H7.5A2.5 2.5 0 0 0 5 5.5v9.258A4.479 4.479 0 0 1 7.5 14h2v2h-2a2.5 2.5 0 0 0 0 5H10v2H7.5A4.5 4.5 0 0 1 3 18.5v-13ZM8 5h2.004v2.004H8V5Zm0 3h2.004v2.004H8V8Zm5 6.615l3.914 3.75l-1.384 1.444L14 18.343V23h-2v-4.657l-1.53 1.466l-1.384-1.445L13 14.615Z"/>`),
		g.Group(children),
	)
}