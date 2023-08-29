package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.62 13.08a.9.9 0 0 0-.54.54a1 1 0 0 0 1.3 1.3a1.15 1.15 0 0 0 .33-.21a1.15 1.15 0 0 0 .21-.33A.84.84 0 0 0 8 14a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21Zm13.14-4L18.4 5.05a3 3 0 0 0-2.84-2H8.44a3 3 0 0 0-2.84 2L4.24 9.11A3 3 0 0 0 2 12v4a3 3 0 0 0 2 2.82V20a1 1 0 0 0 2 0v-1h12v1a1 1 0 0 0 2 0v-1.18A3 3 0 0 0 22 16v-4a3 3 0 0 0-2.24-2.89ZM7.49 5.68A1 1 0 0 1 8.44 5h7.12a1 1 0 0 1 1 .68L17.61 9H6.39ZM20 16a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-3.38-2.92a.9.9 0 0 0-.54.54a1 1 0 0 0 1.3 1.3a.9.9 0 0 0 .54-.54A.84.84 0 0 0 18 14a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21ZM13 13h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}