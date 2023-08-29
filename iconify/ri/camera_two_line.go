package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3.993A1 1 0 0 1 2.992 3h18.016c.548 0 .992.445.992.993v16.014a1 1 0 0 1-.992.993H2.992A.993.993 0 0 1 2 20.007V3.993ZM4 5v14h16V5H4Zm8 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 2a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm5-11h2v2h-2V6Z"/>`),
		g.Group(children),
	)
}