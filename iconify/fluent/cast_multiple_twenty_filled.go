package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastMultipleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3h-3v-1a.5.5 0 0 0-1 0v2a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-3a3.001 3.001 0 0 1 3-3h1a.5.5 0 0 0 0-1H5c-.345 0-.68.044-1 .126V5Zm4.752 5.999a.75.75 0 1 0 0-1.499a.75.75 0 0 0 0 1.499ZM8.5 7a.5.5 0 0 0 0 1a2.502 2.502 0 0 1 2.502 2.502a.5.5 0 1 0 1 0A3.502 3.502 0 0 0 8.5 7Zm0-2.5a.5.5 0 1 0 0 1a5.003 5.003 0 0 1 5.003 5.003a.5.5 0 0 0 1 0A6.003 6.003 0 0 0 8.5 4.5Z"/>`),
		g.Group(children),
	)
}