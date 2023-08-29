package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.372 3.264a.784.784 0 0 0-.252-.658L.252.339V0H6.05l4.482 9.905L14.472 0H20v.339L18.403 1.88a.472.472 0 0 0-.177.452v11.334a.472.472 0 0 0 .177.452l1.56 1.542V16H12.12v-.339l1.615-1.58c.159-.16.159-.207.159-.451V4.468L9.402 15.962h-.606L3.566 4.468v7.704c-.043.324.064.65.29.884l2.101 2.568v.338H0v-.338l2.1-2.568a1.03 1.03 0 0 0 .272-.884V3.264z"/>`),
		g.Group(children),
	)
}