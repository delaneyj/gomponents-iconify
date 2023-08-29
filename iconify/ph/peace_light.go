package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeaceLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm90 102a89.44 89.44 0 0 1-13 46.58l-71-49.7V38.2a90.12 90.12 0 0 1 84 89.8Zm-96-89.8v86.68l-71 49.7A90 90 0 0 1 122 38.2ZM57.92 184.4L122 139.53v78.27a89.93 89.93 0 0 1-64.08-33.4ZM134 217.8v-78.27l64.08 44.87A89.93 89.93 0 0 1 134 217.8Z"/>`),
		g.Group(children),
	)
}