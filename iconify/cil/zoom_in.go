package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M208 96h-32v80H96v32h80v80h32v-80h80v-32h-80V96z"/><path fill="currentColor" d="m479.6 400.4l-81.084-81.084l-62.368-25.767A175.008 175.008 0 0 0 368 192.687c0-97.047-78.953-176-176-176s-176 78.953-176 176s78.953 176 176 176a175.028 175.028 0 0 0 101.619-32.378l25.7 62.2L400.4 479.6a56 56 0 0 0 79.2-79.2ZM48 192.687c0-79.4 64.6-144 144-144s144 64.6 144 144s-64.6 144-144 144s-144-64.599-144-144Zm408.971 264.284a24.029 24.029 0 0 1-33.942 0L346.457 380.4l-23.894-57.835l57.837 23.892l76.573 76.572a24.029 24.029 0 0 1-.002 33.942Z"/>`),
		g.Group(children),
	)
}