package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 4.962A1 1 0 0 1 6.039 4c2.32.089 5.775.883 8.677 3.004C17.664 9.16 20 12.66 20 18a1 1 0 1 1-2 0c0-4.66-1.997-7.577-4.465-9.38C11.021 6.783 7.975 6.077 5.962 6A1 1 0 0 1 5 4.961zm.003 5.967a1 1 0 0 1 1.068-.926c1.314.093 3.257.59 4.899 1.806C12.654 13.055 14 15.049 14 18a1 1 0 1 1-2 0c0-2.248-.987-3.671-2.22-4.584c-1.274-.943-2.831-1.346-3.851-1.418a1 1 0 0 1-.926-1.07zM5 18a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1z"/></g>`),
		g.Group(children),
	)
}