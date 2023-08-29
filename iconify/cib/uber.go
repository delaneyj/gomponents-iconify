package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.958.01C7.744.005.864 6.229.047 14.401h11.115V12a.8.8 0 0 1 .797-.797h8c.443 0 .797.359.797.802V20a.797.797 0 0 1-.797.797h-8a.797.797 0 0 1-.797-.797v-2.401H.047c.885 8.792 8.724 15.198 17.51 14.313c8.792-.88 15.198-8.724 14.313-17.51C31.052 6.23 24.172.011 15.959.011z"/>`),
		g.Group(children),
	)
}