package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngularOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5.5l.137-.48a.5.5 0 0 0-.274 0L7.5.5Zm-7 2l-.137-.48a.5.5 0 0 0-.36.535L.5 2.5Zm1 9l-.497.055a.5.5 0 0 0 .273.392L1.5 11.5Zm6 3l-.224.447a.5.5 0 0 0 .448 0L7.5 14.5Zm6-3l.224.447a.5.5 0 0 0 .273-.392L13.5 11.5Zm1-9l.497.055a.5.5 0 0 0-.36-.536L14.5 2.5Zm-7 .5l.458-.2L7.5 1.753L7.042 2.8L7.5 3ZM7.363.02l-7 2l.274.96l7-2l-.274-.96ZM.003 2.554l1 9l.994-.11l-1-9l-.994.11Zm1.273 9.392l6 3l.448-.894l-6-3l-.448.894Zm6.448 3l6-3l-.448-.894l-6 3l.448.894Zm6.273-3.392l1-9l-.994-.11l-1 9l.994.11Zm.64-9.536l-7-2l-.274.962l7 2l.274-.962ZM4.458 11.2l3.5-8l-.916-.4l-3.5 8l.916.4Zm2.584-8l3.5 8l.916-.4l-3.5-8l-.916.4ZM5 9h5V8H5v1Z"/>`),
		g.Group(children),
	)
}