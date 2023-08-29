package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21 512h107q21 0 21-21q0-22-21-22H73l134-134q13-15 0-30q-15-13-30 0L43 439v-55q0-21-22-21q-21 0-21 21v107q0 9 6 15t15 6zm470-149q-22 0-22 21v55L335 305q-15-13-30 0q-13 15 0 30l134 134h-55q-21 0-21 22q0 21 21 21h107q9 0 15-6t6-15V384q0-21-21-21zm0-363H384q-21 0-21 21q0 22 21 22h55L305 177q-13 15 0 30q6 6 15 6t15-6L469 73v55q0 21 22 21q21 0 21-21V21q0-9-6-15t-15-6zM21 149q22 0 22-21V73l134 134q6 6 15 6t15-6q13-15 0-30L73 43h55q21 0 21-22q0-21-21-21H21Q12 0 6 6T0 21v107q0 21 21 21z"/>`),
		g.Group(children),
	)
}