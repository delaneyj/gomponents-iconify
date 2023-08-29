package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3.993A1 1 0 0 1 2.992 3h18.016c.548 0 .992.445.992.993v16.014a1 1 0 0 1-.992.993H2.992A.993.993 0 0 1 2 20.007V3.993ZM12 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm0 2a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm6-12v2h2V5h-2Z"/>`),
		g.Group(children),
	)
}