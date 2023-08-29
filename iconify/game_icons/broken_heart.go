package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrokenHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M373.47 25.5c-33.475-.064-67.614 13.444-94.44 43.156l37.22 145.156l-33.437.032l35.343 132.093l-116.718-188.375l50.03 5.375L202.5 47.312C120.437-1.43 4.756 40.396 8.5 158.156c4.402 138.44 191.196 184.6 247.406 331.625c59.376-147.035 251.26-184.33 246.656-331.624c-2.564-82.042-64.6-132.532-129.093-132.656z"/>`),
		g.Group(children),
	)
}