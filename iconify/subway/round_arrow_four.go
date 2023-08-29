package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundArrowFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M254.5 157.5h-95.9c24.8-24.2 58.5-39.4 95.9-39.4c69.4 0 126.2 51.4 135.9 118.2h79.8c-10-110.4-102.6-196.9-215.6-196.9c-62.4 0-118.1 26.8-157.5 69V0L37.9 59.1v157.5h157.5l59.1-59.1zm59.1 137.9l-59.1 59.1h95.9c-24.8 24.2-58.5 39.4-95.9 39.4c-69.4 0-126.2-51.4-135.9-118.2H38.9c10 110.4 102.6 196.9 215.6 196.9c62.4 0 118.1-26.8 157.5-69V512l59.1-59.1V295.4H313.6z"/>`),
		g.Group(children),
	)
}