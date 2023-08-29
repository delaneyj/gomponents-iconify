package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M399 15Q386 2 369.5 2T340 15L37 316q-13 13-13 30t13 30l30 32q13 13 29.5 13t29.5-13l303-303q13-13 13-30t-13-30zm-91 152l-30-30l89-92l30 30zM67 3Q45 3 45 24v21H24Q3 45 3 67q0 21 21 21h21v21q0 22 22 22q9 0 15-6t6-16V88h21q10 0 16-6t6-15q0-22-22-22H88V24Q88 3 67 3zm298 256q-21 0-21 21v21h-21q-22 0-22 22q0 9 6 15t16 6h21v21q0 10 6 16t15 6q22 0 22-22v-21h21q21 0 21-21q0-22-21-22h-21v-21q0-21-22-21z"/>`),
		g.Group(children),
	)
}