package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7H3V3m14 0h4v4h-4zm-7 3V3h4v4h-3m-8 3h4v4H3zm14 3v-3h4v4h-3m-4 0h-4v-4m0 7h4v4h-4zM3 3l18 18"/>`),
		g.Group(children),
	)
}