package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 4.25a.75.75 0 0 0 1.5 0v-2.5a.75.75 0 0 0-1.5 0v2.5Zm4.91.97a.75.75 0 0 1 1.06 0l2.25 2.25L16 8l-.53.53l-2.25 2.25a.75.75 0 0 1-1.06-1.06L13.88 8l-1.72-1.72a.75.75 0 0 1 0-1.06ZM8 10a.75.75 0 0 1-.75-.75v-2.5a.75.75 0 0 1 1.5 0v2.5A.75.75 0 0 1 8 10Zm0 5a.75.75 0 0 1-.75-.75v-2.5a.75.75 0 0 1 1.5 0v2.5A.75.75 0 0 1 8 15ZM2.78 5.22a.75.75 0 0 1 1.06 1.06L2.12 8l1.72 1.72a.75.75 0 1 1-1.06 1.06L.53 8.53L0 8l.53-.53l2.25-2.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}