package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M387 128c-11.8 0-21.3 9.5-21.3 21.3v213.3c0 58.9-47.7 106.7-106.7 106.7s-106.7-47.8-106.7-106.7v-256c0-35.3 28.6-64 64-64s64 28.7 64 64V320c0 11.8-9.6 21.3-21.3 21.3s-21.3-9.5-21.3-21.3V149.3c0-11.8-9.6-21.3-21.3-21.3c-11.8 0-21.3 9.5-21.3 21.3V320c0 35.4 28.6 64 64 64s64-28.6 64-64V106.7C323 47.8 275.2 0 216.3 0S109.7 47.8 109.7 106.7v256c0 82.5 66.9 149.3 149.3 149.3s149.3-66.9 149.3-149.3V149.3c0-11.8-9.5-21.3-21.3-21.3z"/>`),
		g.Group(children),
	)
}