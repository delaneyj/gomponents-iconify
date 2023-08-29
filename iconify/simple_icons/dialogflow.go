package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dialogflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.996 0a1.639 1.639 0 0 0-.82.22L3.344 4.74a1.648 1.648 0 0 0-.535.498l9.136 5.28l9.213-5.32a1.652 1.652 0 0 0-.51-.458L12.818.22a1.639 1.639 0 0 0-.822-.22zm9.336 5.5l-9.387 5.422l-9.3-5.373a1.648 1.648 0 0 0-.12.615v9.043a1.643 1.643 0 0 0 .819 1.42l3.918 2.266v4.617a.493.493 0 0 0 .74.424l12.654-7.303a1.639 1.639 0 0 0 .819-1.42V6.162a1.652 1.652 0 0 0-.143-.662z"/>`),
		g.Group(children),
	)
}