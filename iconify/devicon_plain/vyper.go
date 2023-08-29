package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vyper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M.953 9.125L16 35.187L31.047 9.125H.953zm96 0L112 35.187l15.047-26.062H96.953zM32 9.675L16.637 36.288L32 62.897l15.363-26.61L32 9.676zm64 0L80.637 36.288L96 62.897l15.363-26.61L96 9.676zM56.318 22.982l-7.683 13.307L64 62.9l15.363-26.613L71.68 22.98H56.318zM48 37.389L40.29 50.74L32.636 64L48 90.614L63.363 64L48 37.39zm32 .002L64.637 64L80 90.61L95.363 64L80 37.39zM64 65.1L48.637 91.714L64 118.323l15.363-26.61L64 65.102z"/>`),
		g.Group(children),
	)
}