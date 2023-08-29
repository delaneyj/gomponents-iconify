package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M426.072 86.928A238.75 238.75 0 0 0 88.428 424.572A238.75 238.75 0 0 0 426.072 86.928ZM257.25 462.5c-114 0-206.75-92.748-206.75-206.75S143.248 49 257.25 49S464 141.748 464 255.75S371.252 462.5 257.25 462.5Z"/><path fill="currentColor" d="m221.27 305.808l-73.413-73.412l-22.627 22.627l96.04 96.04l167.5-167.499l-22.628-22.627L221.27 305.808z"/>`),
		g.Group(children),
	)
}