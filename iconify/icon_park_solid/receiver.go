package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Receiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSReceiver0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M9.858 38.142c7.81 7.81 20.474 7.81 28.284 0L9.858 9.858c-7.81 7.81-7.81 20.474 0 28.284Z"/><path d="m33.9 33.9l5.27-21.986M24 24l13.172-13.172M14.1 14.1l21.986-5.27"/><path fill="#fff" d="M44 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSReceiver0)"/>`),
		g.Group(children),
	)
}