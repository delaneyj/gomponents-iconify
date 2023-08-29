package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RetroBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M35 14C35 8.47715 30.0751 4 24 4C17.9249 4 13 8.47715 13 14"/><path fill="#2F88FF" d="M7 16C7 14.8954 7.89543 14 9 14H39C40.1046 14 41 14.8954 41 16V28C41 29.1046 40.1046 30 39 30H30C28.8954 30 28 29.1046 28 28V28C28 26.8954 27.1046 26 26 26H22C20.8954 26 20 26.8954 20 28V28C20 29.1046 19.1046 30 18 30H9C7.89543 30 7 29.1046 7 28V16Z"/><path d="M10 30V44H38V30"/><rect width="8" height="6" x="20" y="26"/></g>`),
		g.Group(children),
	)
}