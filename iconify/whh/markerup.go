package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markerup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M384 1024q-104 0-192.5-51.5t-140-140T0 640q0-125 73-225t188-139L384 0l123 276q115 39 188 139t73 225q0 104-51.5 192.5t-140 140T384 1024zm0-640q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75z"/>`),
		g.Group(children),
	)
}