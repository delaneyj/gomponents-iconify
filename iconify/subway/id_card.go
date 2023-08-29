package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 394c0 23.5 19.1 42.7 42.7 42.7h426.7c23.5 0 42.7-19.1 42.7-42.7v-42.7H0V394zM469.3 95.3H42.7C19.1 95.3 0 114.5 0 138v170.7h512V138c0-23.5-19.1-42.7-42.7-42.7z"/>`),
		g.Group(children),
	)
}