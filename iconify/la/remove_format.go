package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4.29 5.973L3.083 7.57l24.693 18.608l1.203-1.598l-10.953-8.254L20.285 10H25v1.5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-16a.5.5 0 0 0-.5.5v2.777L4.29 5.973zM12 10h5.285l-1.623 4.545l-3.728-2.81A.488.488 0 0 0 12 11.5V10zm2.254 8.49L13 22h-1.5a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5H16l.62-1.734l-2.366-1.776z"/>`),
		g.Group(children),
	)
}