package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfTiltLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200.12 55.88A102 102 0 1 0 55.88 200.13A102 102 0 1 0 200.12 55.88ZM90 209.62a89.61 89.61 0 0 1-21.23-13.89L90 174.49Zm32 8.16a90 90 0 0 1-20-3.58v-51.71l20-20Zm32-3.58a89.8 89.8 0 0 1-20 3.58v-87.29l20-20Zm32-17.4a89.45 89.45 0 0 1-20 12.83V98.49l20-20Zm-125.73-9.56a90 90 0 0 1 127-127ZM198 184.57V71.43a90 90 0 0 1 0 113.14Z"/>`),
		g.Group(children),
	)
}