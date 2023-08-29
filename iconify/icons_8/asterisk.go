package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m17.7 17l6.2 8.4l-2.4 1.6l-5.5-8.7l-5.5 8.7l-2.3-1.6l6.2-8.4l-9.3-2.4L6 12l9.1 3.2L14.5 5h3L17 15.2l9-3.2l.8 2.7l-9.1 2.3z"/>`),
		g.Group(children),
	)
}