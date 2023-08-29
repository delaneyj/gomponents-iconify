package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAside(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 1h18v12l-6 6H1V1zm3 3v1h12V4H4zm0 4v1h12V8H4zm6 5v-1H4v1h6zm2 4l5-5h-5v5z"/>`),
		g.Group(children),
	)
}