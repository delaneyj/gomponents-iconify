package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 19V5h6v14h-6Zm-8 0V5h6v14H5Zm10-2h2V7h-2v10Zm-8 0h2V7H7v10ZM7 7v10V7Zm8 0v10V7Z"/>`),
		g.Group(children),
	)
}