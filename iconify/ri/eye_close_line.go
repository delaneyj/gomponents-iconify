package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeCloseLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.343 18.782l-1.932-.518l.787-2.939a10.99 10.99 0 0 1-3.237-1.872l-2.153 2.154l-1.414-1.414l2.153-2.154a10.957 10.957 0 0 1-2.371-5.07l1.968-.359a9.002 9.002 0 0 0 17.713 0l1.968.358a10.958 10.958 0 0 1-2.372 5.071l2.154 2.154l-1.414 1.414l-2.154-2.154a10.991 10.991 0 0 1-3.237 1.872l.788 2.94l-1.932.517l-.788-2.94a11.068 11.068 0 0 1-3.74 0l-.787 2.94Z"/>`),
		g.Group(children),
	)
}