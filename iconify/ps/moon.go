package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M299 235q-63 0-106.5-43.5T149 85q0-40 26-85q-74 7-124.5 62T0 192q0 80 56 136t136 56q75 0 130-50.5T384 209q-45 26-85 26zM192 341q-62 0-105.5-43.5T43 192q0-79 64-122v15q0 80 56 136t136 56h15q-43 64-122 64z"/>`),
		g.Group(children),
	)
}