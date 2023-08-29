package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListNumberedRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M40 80h288v32H40zm0 160h288v32H40zm0 160h288v32H40zm400-240V40h-64v32h32v88h32zm-64 102.111V312h80v-32h-44.223L456 257.889V192h-80v32h48v14.111l-48 24zM376 440v32h80V344h-80v32h48v16h-24v32h24v16h-48z"/>`),
		g.Group(children),
	)
}