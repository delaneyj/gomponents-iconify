package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionBearRightTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2ZM6 7.414L7.414 6L14 12.586L12.586 14ZM26 16h-2V9.414l-6.536 6.536A4.968 4.968 0 0 0 16 19.485V26h-2v-6.515a6.954 6.954 0 0 1 2.05-4.95L22.586 8H16V6h10Z"/><path fill="none" d="M26 6v10h-2V9.414l-6.536 6.536A4.968 4.968 0 0 0 16 19.485V26h-2v-6.515a6.954 6.954 0 0 1 2.05-4.95L22.586 8H16V6Zm-12 6.586L7.414 6L6 7.414L12.586 14Z"/>`),
		g.Group(children),
	)
}