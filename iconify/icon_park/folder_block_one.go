package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderBlockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M43 23V14C43 12.8954 42.1046 12 41 12H24L19 6H7C5.89543 6 5 6.89543 5 8V40C5 41.1046 5.89543 42 7 42H22"/><circle cx="35" cy="35" r="8" fill="#2F88FF" stroke="#000"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M37 33L33 37"/></g>`),
		g.Group(children),
	)
}