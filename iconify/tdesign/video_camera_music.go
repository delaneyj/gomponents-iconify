package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 4.234v15.49l-7-4V20H0V4h17v4.434l7-4.2Zm-7 6.532v2.654l5 2.857v-8.51l-5 3ZM15 6H2v12h13V6Zm-2.5 3h-2v5a3 3 0 1 1-2-2.83V7h4v2Zm-4 5a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/>`),
		g.Group(children),
	)
}