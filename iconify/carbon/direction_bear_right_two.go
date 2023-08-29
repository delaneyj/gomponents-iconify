package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionBearRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 4v2h6.586l-6.536 6.536a6.954 6.954 0 0 0-2.05 4.95V28h2V17.485a4.968 4.968 0 0 1 1.464-3.535L24 7.414V14h2V4ZM6 7.414L7.414 6L13 11.586L11.586 13z"/>`),
		g.Group(children),
	)
}