package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Character(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M255 45.4c-24.5 0-47 11.8-63.9 33.4c-16.9 21.5-27.1 52.6-27.1 86.5c0 36 12.1 67.5 31 89.5l13.5 15l-19.6 4.6c-52.3 11.9-77.4 36.9-91.75 75.2c-13.7 35.7-15.6 84.8-16.1 143.3H431c-.2-58.7-.5-109.3-13-145.5c-13.4-39.4-37.9-64.3-94-75.4l-19.9-3.7l12.9-15.7c17.7-21.9 28.8-52.6 28.8-87.5c0-33.9-10.3-64.9-27.2-86.3c-16.8-21.7-39.3-33.6-63.6-33.4z"/>`),
		g.Group(children),
	)
}