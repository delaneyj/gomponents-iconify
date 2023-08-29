package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29 19.41a1 1 0 0 1-.71-.29L18 8.83L7.71 19.12a1 1 0 0 1-1.41-1.41L18 6l11.71 11.71a1 1 0 0 1-.71 1.7Z" class="clr-i-outline clr-i-outline-path-1"/><path fill="currentColor" d="M29 30.41a1 1 0 0 1-.71-.29L18 19.83L7.71 30.12a1 1 0 0 1-1.41-1.41L18 17l11.71 11.71a1 1 0 0 1-.71 1.7Z" class="clr-i-outline clr-i-outline-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}