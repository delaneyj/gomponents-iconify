package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 4.234v15.49l-7-4V20H0V4h17v4.434l7-4.2Zm-7 6.532v2.654l5 2.857v-8.51l-5 3ZM15 18V6H2v12h13Zm-2.5-5h-8v-2h8v2Z"/>`),
		g.Group(children),
	)
}