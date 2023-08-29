package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceSixFacesThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M74.5 36A38.5 38.5 0 0 0 36 74.5v363A38.5 38.5 0 0 0 74.5 476h363a38.5 38.5 0 0 0 38.5-38.5v-363A38.5 38.5 0 0 0 437.5 36h-363zm316.97 36.03A50 50 0 0 1 440 122a50 50 0 0 1-100 0a50 50 0 0 1 51.47-49.97zM256 206a50 50 0 0 1 0 100a50 50 0 0 1 0-100zM123.47 340.03A50 50 0 0 1 172 390a50 50 0 0 1-100 0a50 50 0 0 1 51.47-49.97z"/>`),
		g.Group(children),
	)
}