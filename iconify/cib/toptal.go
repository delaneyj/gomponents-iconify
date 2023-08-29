package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toptal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.969 13.385L13.584 0l-2.719 2.719l5.031 5.026l-10.875 10.87L18.412 32l2.719-2.719l-5.031-5.026l10.88-10.87zm-15.901 5.641l8.089-8.083l1.63 1.63l-8.083 8.151l-1.63-1.682z"/>`),
		g.Group(children),
	)
}