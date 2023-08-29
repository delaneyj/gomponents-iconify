package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Key(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M373.3 0c-76.6 0-138.7 62.1-138.7 138.7c0 16 2.8 31.2 7.8 45.5L21.3 405.3L0 512h128v-42.7l21.3-21.3l21.3 21.3l42.7-42.7l-21.3-21.3l42.7-42.7L256 384l42.7-42.7l-21.4-21.3l50.5-50.5c14.3 5 29.6 7.8 45.5 7.8c76.6 0 138.7-62.1 138.7-138.7S449.9 0 373.3 0zm32 149.3c-23.5 0-42.7-19.1-42.7-42.7S381.8 64 405.3 64c23.5 0 42.7 19.1 42.7 42.7s-19.1 42.6-42.7 42.6z"/>`),
		g.Group(children),
	)
}