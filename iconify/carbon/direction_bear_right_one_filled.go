package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionBearRightOneFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm-6 14h-2V9.414l-6.536 6.536A4.968 4.968 0 0 0 12 19.485V26h-2v-6.515a6.954 6.954 0 0 1 2.05-4.95L18.586 8H12V6h10Z"/><path fill="none" d="M22 16h-2V9.414l-6.536 6.536A4.968 4.968 0 0 0 12 19.485V26h-2v-6.515a6.954 6.954 0 0 1 2.05-4.95L18.586 8H12V6h10Z"/>`),
		g.Group(children),
	)
}