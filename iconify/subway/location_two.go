package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M264 0C157.3 0 72 85.3 72 192c0 36.9 11 65.4 30.1 94.3l141.7 215c4.3 6.5 11.7 10.7 20.2 10.7s16-4.3 20.2-10.7l141.7-215C445 257.4 456 228.9 456 192C456 85.3 370.7 0 264 0zm0 341.2c-82.5 0-149.3-66.9-149.3-149.5c0-82.5 66.9-149.5 149.3-149.5c82.5 0 149.3 66.9 149.3 149.5S346.5 341.2 264 341.2zm64-170.5h-42.7v-64c0-11.8-9.5-21.3-21.3-21.3s-21.3 9.5-21.3 21.3V192c0 11.8 9.6 21.3 21.3 21.3h64c11.8 0 21.3-9.5 21.3-21.3s-9.5-21.3-21.3-21.3z"/>`),
		g.Group(children),
	)
}