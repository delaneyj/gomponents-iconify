package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetlifyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m.5 7.5l-.354-.354a.5.5 0 0 0 0 .708L.5 7.5Zm7-7l.354-.354a.5.5 0 0 0-.708 0L7.5.5Zm7 7l.354.354a.5.5 0 0 0 0-.708L14.5 7.5Zm-7 7l-.354.354a.5.5 0 0 0 .708 0L7.5 14.5ZM.854 7.854l7-7l-.708-.708l-7 7l.708.708Zm6.292-7l7 7l.708-.708l-7-7l-.708.708Zm7 6.292l-7 7l.708.708l7-7l-.708-.708Zm-6.292 7l-7-7l-.708.708l7 7l.708-.708ZM4.314 3.964l10 4l.372-.928l-10-4l-.372.928ZM8.58 1.728l-5.5 8.5l.84.544l5.5-8.5l-.84-.544ZM2.1 5.8l6 8l.8-.6l-6-8l-.8.6ZM.394 7.989l11.5 2.5l.212-.978l-11.5-2.5l-.212.978Zm2.32 1.963l9.5-4.5l-.428-.904l-9.5 4.5l.428.904Zm7.292-6.53l-1.5 9.5l.988.156l1.5-9.5l-.988-.156Z"/>`),
		g.Group(children),
	)
}