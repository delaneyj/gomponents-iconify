package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.82 4.106c-.234-.417-.671-.601-1.06-.606c-.39-.005-.838.17-1.068.6l-.164.307a77.416 77.416 0 0 0-5.986 14.56c-.158.537.141.984.498 1.199c.356.213.86.264 1.291-.008l4.632-2.925c.574-.363 1.388-.358 1.99.029l4.727 3.04c.425.273.927.244 1.29.031c.37-.217.652-.673.479-1.207a82.268 82.268 0 0 0-6.455-14.708l-.174-.312Z"/>`),
		g.Group(children),
	)
}