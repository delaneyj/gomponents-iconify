package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Elevator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800.857 256h-32v736q0 13-9.5 22.5t-22.5 9.5h-640q-13 0-22.5-9.5t-9.5-22.5V256h-32q-13 0-22.5-9.5T.857 224V32q0-13 9.5-22.5t22.5-9.5h768q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm-416 704h64V256h-64v704zm32-896q-66 0-113 37.5t-47 90.5h320q0-53-47-90.5t-113-37.5z"/>`),
		g.Group(children),
	)
}