package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunglassesLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 42a6 6 0 0 0 0 12a18 18 0 0 1 18 18v58H38V72a18 18 0 0 1 18-18a6 6 0 0 0 0-12a30 30 0 0 0-30 30v92a42 42 0 0 0 84 0v-22h36v22a42 42 0 0 0 84 0V72a30 30 0 0 0-30-30ZM38 164v-21.52L84.53 189A30 30 0 0 1 38 164Zm60 0a29.83 29.83 0 0 1-5 16.53L54.48 142H98Zm60 0v-21.52L204.53 189A30 30 0 0 1 158 164Zm55 16.53L174.48 142H218v22a29.83 29.83 0 0 1-5 16.53Z"/>`),
		g.Group(children),
	)
}