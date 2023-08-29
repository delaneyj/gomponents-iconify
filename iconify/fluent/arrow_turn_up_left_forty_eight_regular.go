package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnUpLeftFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M35.5 40.75a1.25 1.25 0 1 0 2.5 0v-18.5A7.25 7.25 0 0 0 30.75 15H12.582l7.04-6.854a1.25 1.25 0 0 0-1.744-1.792l-9.5 9.25a1.25 1.25 0 0 0 0 1.792l9.5 9.25a1.25 1.25 0 0 0 1.744-1.792L12.069 17.5H30.75a4.75 4.75 0 0 1 4.75 4.75v18.5Z"/>`),
		g.Group(children),
	)
}