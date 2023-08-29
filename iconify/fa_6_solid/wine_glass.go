package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineGlass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32.1 29.3C33.5 12.8 47.4 0 64 0h192c16.6 0 30.5 12.8 31.9 29.3l14 168.4c6 72-42.5 135.2-109.9 150.6V448h48c17.7 0 32 14.3 32 32s-14.3 32-32 32H80c-17.7 0-32-14.3-32-32s14.3-32 32-32h48v-99.6C60.6 333 12.1 269.8 18.1 197.8l14-168.4zm56 98.7h143.8l-5.3-64H93.4l-5.3 64z"/>`),
		g.Group(children),
	)
}