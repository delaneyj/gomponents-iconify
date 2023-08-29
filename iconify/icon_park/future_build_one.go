package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FutureBuildOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M13.9994 24C15.2168 15.7039 23.9994 5 23.9994 5C23.9994 5 32.7819 15.7039 33.9994 24C35.09 31.4323 30.9994 44 30.9994 44H16.9994C16.9994 44 12.9087 31.4323 13.9994 24Z"/><path d="M18 14H30"/><path d="M15 20H33"/><path d="M14 26L34 26"/><path d="M15 32H33"/><path d="M16 38H32"/><path stroke-linejoin="round" d="M4 44H44"/><path d="M24 4V6"/></g>`),
		g.Group(children),
	)
}