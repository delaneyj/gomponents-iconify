package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slideronefull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.743 576h-640V448h640q27 0 45.5 19t18.5 45.5t-18.5 45t-45.5 18.5zm32-320h-192q-13 0-22.5-9.5t-9.5-22.5V32q0-13 9.5-22.5t22.5-9.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5zm-992-127.5q0-26.5 19-45.5t45-19h640v128h-640q-26 0-45-18.5t-19-45zm32 255.5h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5h-192q-13 0-22.5-9.5T.743 608V416q0-13 9.5-22.5t22.5-9.5zm0 384h192q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5t-22.5 9.5h-192q-13 0-22.5-9.5T.743 992V800q0-13 9.5-22.5t22.5-9.5zm992 128.5q0 26.5-18.5 45t-45.5 18.5h-640V832h640q27 0 45.5 19t18.5 45.5z"/>`),
		g.Group(children),
	)
}