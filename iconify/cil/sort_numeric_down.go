package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortNumericDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M151.196 16v417.568l-51.883-51.881l-22.626 22.626l90.509 90.51l90.51-90.51l-22.627-22.627l-51.883 51.882V16h-32zM432 200h-40V56h-32v32h-32v32h32v80h-40v32h112v-32zm-76.8 232H336v32h19.2a76.887 76.887 0 0 0 76.8-76.8v-51.6a55.663 55.663 0 0 0-55.6-55.6H372a55.663 55.663 0 0 0-55.6 55.6v4.4a55.663 55.663 0 0 0 55.6 55.6h4.4a55.262 55.262 0 0 0 23.474-5.215A44.849 44.849 0 0 1 355.2 432Zm21.2-68.4H372a23.627 23.627 0 0 1-23.6-23.6v-4.4A23.627 23.627 0 0 1 372 312h4.4a23.627 23.627 0 0 1 23.6 23.6v4.4a23.627 23.627 0 0 1-23.6 23.6Z"/>`),
		g.Group(children),
	)
}