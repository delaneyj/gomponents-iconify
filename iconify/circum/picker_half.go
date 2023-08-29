package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PickerHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.936 5.889a2.825 2.825 0 0 0-4.81-2.02l-2.21 2.22l-.75-.75a.771.771 0 0 0-.55-.22a.8.8 0 0 0-.55.22a.785.785 0 0 0 0 1.1l.75.75l-8.76 8.76a3.154 3.154 0 0 0-.92 2.13l-.07 1.52a1.316 1.316 0 0 0 1.28 1.35h.06l1.52-.07a3.21 3.21 0 0 0 2.13-.93l8.76-8.76l.75.75a.8.8 0 0 0 1.1 0a.785.785 0 0 0 0-1.1l-.75-.75l2.18-2.18a2.845 2.845 0 0 0 .84-2.02Zm-8.56 8.33H7.2l6.33-6.32l2.59 2.59ZM19.4 7.2l-2.18 2.19l-2.594-2.59l2.21-2.22a1.823 1.823 0 0 1 2.56 0a1.859 1.859 0 0 1 0 2.62Z"/>`),
		g.Group(children),
	)
}