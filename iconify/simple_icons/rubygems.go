package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rubygems(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.81 7.9l-2.97 2.95l7.19 7.18l2.96-2.95l4.22-4.23l-2.96-2.96v-.01H7.8zM12 0L1.53 6v12L12 24l10.47-6V6L12 0zm8.47 16.85L12 21.73l-8.47-4.88V7.12L12 2.24l8.47 4.88v9.73z"/>`),
		g.Group(children),
	)
}