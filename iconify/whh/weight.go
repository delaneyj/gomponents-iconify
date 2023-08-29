package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M942 1024H82q-31 0-50-13T6.5 978.5t-6-45t5.5-47T18 844l90-490q14-47 55-72.5t89-25.5h80q-12-33-12-64q0-80 56-136T512 0t136 56t56 136q0 31-12 64h80q48 0 89 25.5t55 72.5l90 490q7 21 12 42.5t5.5 47t-6 45T992 1011t-50 13zM512 128q-27 0-45.5 18.5t-18.5 45t18.5 45.5t45.5 19t45.5-19t18.5-45.5t-18.5-45T512 128zm96 640h-32q-24 0-44-21t-20-43v-64h32q13 0 22.5-9.5T576 608t-9.5-22.5T544 576h-32v-96q0-13-9.5-22.5T480 448t-22.5 9.5T448 480v96h-32q-13 0-22.5 9.5T384 608t9.5 22.5T416 640h32v64q0 48 40 88t88 40h32q13 0 22.5-9.5T640 800t-9.5-22.5T608 768z"/>`),
		g.Group(children),
	)
}