package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.6 15.18a1 1 0 0 0-1.42 1.42l1.06 1.06a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm-3.79-5.4l-1.06 1.06a.91.91 0 0 0-.19.26a1 1 0 0 0-.27 1.61l1.06 1.06a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L4.41 12l.81-.81a1 1 0 0 0-1.41-1.41Zm4.19.79a1 1 0 0 0-1 1v1.5a1 1 0 0 0 2 0v-1.5a1 1 0 0 0-1-1Zm13.71.72l-5-5a1 1 0 0 0-1.09-.21A1 1 0 0 0 15 7v10a1 1 0 0 0 .62.92a.84.84 0 0 0 .38.08a1 1 0 0 0 .71-.29l5-5a1 1 0 0 0 0-1.42ZM17 14.59V9.41L19.59 12ZM12 2a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1ZM8.38 6.08a1 1 0 0 0-1.09.21L6.64 7a1 1 0 0 0 0 1.41a1 1 0 0 0 .7.3a1 1 0 0 0 .45-.11A1 1 0 0 0 9 7.57V7a1 1 0 0 0-.62-.92Z"/>`),
		g.Group(children),
	)
}