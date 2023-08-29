package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M59.1 236.3h393.8s-19.7-98.5-98.5-98.5V59.1L413.5 0h-315l59.1 59.1v78.8c-78.8-.1-98.5 98.4-98.5 98.4zM256 512l39.4-256h-78.8L256 512z"/>`),
		g.Group(children),
	)
}