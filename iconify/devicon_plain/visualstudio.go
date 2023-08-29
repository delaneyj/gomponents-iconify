package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Visualstudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="m95 2.3l30.5 12.3v98.7l-30.7 12.4l-49-48.7l-31 24.1l-12.3-6.2V33.1l12.3-5.9l31 24.3zM14.8 45.7v37.5l18.5-19zm48.1 18.5l31.9 25.1V39z"/>`),
		g.Group(children),
	)
}