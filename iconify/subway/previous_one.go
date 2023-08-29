package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviousOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472.8 90.1L298.6 192.5v-73.9c0-28.4-17.6-41.2-39.2-28.5l-152.8 89.8v-52c0-23.6-19.1-42.7-42.7-42.7H42.7C19.1 85.3 0 104.4 0 128v256c0 23.5 19.1 42.7 42.7 42.7H64c23.5 0 42.7-19.1 42.7-42.7v-51.9l152.8 89.8c21.5 12.7 39.2-.2 39.2-28.5v-74l174.2 102.5c21.5 12.7 39.2-.2 39.2-28.5V118.6c-.1-28.4-17.7-41.2-39.3-28.5z"/>`),
		g.Group(children),
	)
}