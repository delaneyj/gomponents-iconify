package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 14.86v-.05a2.61 2.61 0 0 0-.1-.57l-1.64-9.73a3 3 0 0 0-3-2.51H6.69a3 3 0 0 0-2.95 2.51l-1.62 9.71a2.61 2.61 0 0 0-.1.57v.05C2 14.91 2 15 2 15v4a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4.14ZM5.71 4.83a1 1 0 0 1 1-.83h10.6a1 1 0 0 1 1 .83l1.2 7.22A2.63 2.63 0 0 0 19 12H5a2.63 2.63 0 0 0-.49.05ZM20 19a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-3.92l.08-.46A1 1 0 0 1 5 14h14a1 1 0 0 1 .92.62l.08.46Zm-3-3a1 1 0 1 0 1 1a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}