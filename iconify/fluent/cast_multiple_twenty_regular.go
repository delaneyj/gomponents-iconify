package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastMultipleTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 5a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3h-4v1a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-3c0-1.306.835-2.417 2-2.83V5Zm6 7.5a.5.5 0 0 1 1 0v.5h4a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v4h1a.5.5 0 0 1 0 1H5a2 2 0 0 0-2 2v3a2 2 0 0 0 2 2h3a2 2 0 0 0 2-2v-2.5Zm-.498-2.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM8 7.5a.5.5 0 0 1 .5-.5a3.502 3.502 0 0 1 3.502 3.502a.5.5 0 1 1-1 0A2.502 2.502 0 0 0 8.5 8a.5.5 0 0 1-.5-.5ZM8 5a.5.5 0 0 1 .5-.5a6.003 6.003 0 0 1 6.003 6.003a.5.5 0 0 1-1 0A5.003 5.003 0 0 0 8.5 5.5A.5.5 0 0 1 8 5Z"/>`),
		g.Group(children),
	)
}