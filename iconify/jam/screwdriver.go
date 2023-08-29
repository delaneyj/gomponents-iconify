package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screwdriver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.343 18.071a1 1 0 0 1 0 1.414L3.93 20.9a1 1 0 0 1-1.414 0L1.1 19.485a1 1 0 0 1 0-1.414l1.414-1.414a1 1 0 0 1 1.414 0L9.586 11a2 2 0 0 1 0-2.828l7.07-7.071a2 2 0 0 1 2.83 0l1.413 1.414a2 2 0 0 1 0 2.828l-7.07 7.071a2 2 0 0 1-2.829 0l-5.657 5.657zM12.414 11l7.071-7.071l-1.414-1.414L11 9.585L12.414 11z"/>`),
		g.Group(children),
	)
}