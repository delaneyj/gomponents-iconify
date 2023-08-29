package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90ZM110 96v64a6 6 0 0 1-12 0V96a6 6 0 0 1 12 0Zm48 0v64a6 6 0 0 1-12 0V96a6 6 0 0 1 12 0Z"/>`),
		g.Group(children),
	)
}