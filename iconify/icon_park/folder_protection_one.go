package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderProtectionOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M43 23V14C43 12.8954 42.1046 12 41 12H24L19 6H7C5.89543 6 5 6.89543 5 8V40C5 41.1046 5.89543 42 7 42H22"/><path fill="#2F88FF" d="M28 30.8C28 29.8667 34 28 34 28C34 28 40 29.8667 40 30.8C40 38.2667 34 42 34 42C34 42 28 38.2667 28 30.8Z"/></g>`),
		g.Group(children),
	)
}