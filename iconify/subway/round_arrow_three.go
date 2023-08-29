package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundArrowThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M247.6 393.8c-37.4 0-71.1-15.2-95.9-39.4h95.9l-59.1-59.1H31v157.5L90.1 512V403.6c39.5 42.2 95.2 69 157.5 69c113 0 205.7-86.5 215.6-196.9h-79.8c-9.6 66.7-66.4 118.1-135.8 118.1zM405.2 0v108.4c-39.5-42.2-95.2-69-157.5-69C134.6 39.4 42 125.9 32 236.3h79.8c9.6-66.7 66.5-118.2 135.9-118.2c37.4 0 71.1 15.2 95.9 39.4h-95.9l59.1 59.1h157.5V59.1L405.2 0z"/>`),
		g.Group(children),
	)
}