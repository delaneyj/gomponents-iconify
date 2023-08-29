package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M40 472h432V312h-32v128H72V312H40v160zm0-206.245l49.373 25.437l53.82-46.829l56.159 50.528L256 247.052l56.648 47.839l56.159-50.528l53.82 46.829L472 265.755V40H40ZM72 72h368v174.244l-12.738 6.564l-58.809-51.171l-56.471 50.806L256 205.167l-55.982 47.276l-56.471-50.806l-58.809 51.171L72 246.244Z"/>`),
		g.Group(children),
	)
}