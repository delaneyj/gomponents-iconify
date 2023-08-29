package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a1 1 0 0 0-1 1v.1A5.002 5.002 0 0 0 3 7v4.94l-1.78 1.78a.75.75 0 1 0 1.06 1.06L14.776 2.284a.75.75 0 0 0-1.06-1.06l-2.211 2.21A4.991 4.991 0 0 0 9 2.1V2a1 1 0 0 0-1-1Zm0 2.5c.95 0 1.813.379 2.444.995L4.5 10.439V7A3.5 3.5 0 0 1 8 3.5Zm5 4.25a.75.75 0 0 0-1.5 0v3.817l.194.214l.65.719H6.75a.75.75 0 0 0-.728.932l.011.043A2.02 2.02 0 0 0 7.993 15c.737 0 1.389-.4 1.738-1h3.74c.868 0 1.324-1.028.742-1.671L13 10.989V7.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}