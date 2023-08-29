package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstagramOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#2F88FF" stroke="#000" stroke-width="4"><circle cx="9" cy="8" r="4"/><rect width="8" height="25" x="5" y="18" stroke-linejoin="round"/><path stroke-linejoin="round" d="M21 27.5V43H28V29C28 26.5 29.5 24.5 32 24.5C34.5 24.5 36 27 36 29V43H43V27.5C43 24.5 39.5 18 32 18C24.5 18 21 24.5 21 27.5Z"/></g>`),
		g.Group(children),
	)
}