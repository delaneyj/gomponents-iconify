package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Park(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 19.56v4.38M4.5 12.56h3.724A33.264 33.264 0 0 1 1.5 19.31v.25h21v-.25a33.263 33.263 0 0 1-6.724-6.75H19.5v-.25l-1.386-1.04A15.383 15.383 0 0 1 12 .06a15.384 15.384 0 0 1-6.114 11.21L4.5 12.31v.25Z"/>`),
		g.Group(children),
	)
}