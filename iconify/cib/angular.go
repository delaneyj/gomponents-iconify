package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.24 16.859h5.51l-2.755-6.542zM15.995.01L.917 5.317l2.297 19.677l12.781 6.995l12.786-6.984l2.297-19.688L15.995.015zm9.411 24.396H21.89l-1.896-4.667h-8l-1.896 4.667H6.582l9.411-20.865z"/>`),
		g.Group(children),
	)
}