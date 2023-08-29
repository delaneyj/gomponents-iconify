package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiPeople(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22V8.775q-2.275-.6-3.638-2.513T4 2h2q0 2.075 1.338 3.538T10.75 7h2.5q.75 0 1.4.275t1.175.8L20.35 12.6l-1.4 1.4L15 10.05V22h-2v-6h-2v6H9Zm3-16q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Z"/>`),
		g.Group(children),
	)
}