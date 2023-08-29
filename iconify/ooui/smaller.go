package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smaller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 16h-1.05a.44.44 0 0 1-.29-.09a.58.58 0 0 1-.17-.22l-.7-1.84H6.2l-.7 1.84a.56.56 0 0 1-.16.21a.43.43 0 0 1-.29.1H4l3.31-8.35h1.38zm-2.57-3.13L8.28 9.82a8.5 8.5 0 0 1-.28-.9q-.06.27-.14.5l-.14.4l-1.15 3zM15 6l3-4h-6z"/>`),
		g.Group(children),
	)
}