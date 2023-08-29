package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumbered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M184 80h288v32H184zm0 160h288v32H184zm0 160h288v32H184zm-64-240V40H56v32h32v88h32zM56 262.111V312h80v-32H91.777L136 257.889V192H56v32h48v14.111l-48 24zM56 440v32h80V344H56v32h48v16H80v32h24v16H56z"/>`),
		g.Group(children),
	)
}