package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveXCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 8.464L18.965 1.5l.177.177l.152.228a10.097 10.097 0 0 0 2.8 2.801l.23.153l.176.177L15.536 12l6.964 6.965l-.177.176l-.228.153a10.096 10.096 0 0 0-2.801 2.8l-.153.23l-.176.176L12 15.536L5.036 22.5l-.177-.177l-.153-.228a10.098 10.098 0 0 0-2.8-2.801l-.23-.153l-.176-.176L8.465 12L1.5 5.036l.177-.177l.229-.153a10.098 10.098 0 0 0 2.8-2.8l.153-.23l.177-.176L12 8.464Z"/>`),
		g.Group(children),
	)
}