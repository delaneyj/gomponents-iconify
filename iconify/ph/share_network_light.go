package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareNetworkLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M176 162a37.91 37.91 0 0 0-28.3 12.67l-48.9-31.43a37.89 37.89 0 0 0 0-30.48l48.9-31.43a38 38 0 1 0-6.5-10.09l-48.9 31.43a38 38 0 1 0 0 50.66l48.9 31.43A38 38 0 1 0 176 162Zm0-132a26 26 0 1 1-26 26a26 26 0 0 1 26-26ZM64 154a26 26 0 1 1 26-26a26 26 0 0 1-26 26Zm112 72a26 26 0 1 1 26-26a26 26 0 0 1-26 26Z"/>`),
		g.Group(children),
	)
}