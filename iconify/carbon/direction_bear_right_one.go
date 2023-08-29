package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionBearRightOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 4v2h6.586l-6.536 6.536a6.954 6.954 0 0 0-2.05 4.95V28h2V17.485a4.968 4.968 0 0 1 1.464-3.535L20 7.414V14h2V4Z"/>`),
		g.Group(children),
	)
}