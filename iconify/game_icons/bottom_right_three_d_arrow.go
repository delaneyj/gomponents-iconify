package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomRightThreeDArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m126.47 18.25l273.81 229.344l39.314-39.313l23.47 205.095l-205.095-23.438l37.467-37.468L20.594 58.655v99.28l195.25 235.126l5.437 6.532l-6.03 6.03l-45.97 45.97l323.033 38.344l-38.375-323l-48.313 48.312l-6 6l-6.563-5.438L155.032 18.25h-28.563z"/>`),
		g.Group(children),
	)
}