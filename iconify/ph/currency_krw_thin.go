package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyKrwThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 132h-28.06l23.77-58.49a4 4 0 1 0-7.42-3l-25 61.51h-46.6l-25-61.51a4 4 0 0 0-7.42 0L99.31 132H52.69l-25-61.51a4 4 0 0 0-7.42 3L44.06 132H16a4 4 0 0 0 0 8h31.31l25 61.51a4 4 0 0 0 7.42 0l25-61.51h46.62l25 61.51a4 4 0 0 0 7.42 0l25-61.51H240a4 4 0 0 0 0-8ZM76 189.37L55.94 140h40.12ZM107.94 132L128 82.63L148.06 132ZM180 189.37L159.94 140h40.12Z"/>`),
		g.Group(children),
	)
}