package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.346 1.633a.75.75 0 0 1 1.308.735L12.86 12.672l2.708 4.815a3.751 3.751 0 1 1 .136 3.303L12 14.202L8.296 20.79a3.75 3.75 0 1 1 .136-3.303l2.708-4.815L5.346 2.368a.75.75 0 0 1 1.308-.735L12 11.142l5.346-9.51ZM5 21.25a2.25 2.25 0 0 0 1.953-1.131l.02-.036A2.25 2.25 0 1 0 5 21.25Zm14 0a2.25 2.25 0 0 1-1.953-1.131l-.02-.036A2.25 2.25 0 1 1 19 21.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}