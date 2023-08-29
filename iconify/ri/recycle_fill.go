package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecycleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.561 12.098l1.532 2.652A3.5 3.5 0 0 1 18.06 20H16v2l-5-3.5l5-3.5v2h2.062a.5.5 0 0 0 .471-.668l-.038-.082l-1.53-2.652l2.597-1.5ZM7.736 9.384l.531 6.08l-1.731-1l-1.032 1.786a.5.5 0 0 0 .343.742l.09.008H9v3H5.937a3.5 3.5 0 0 1-3.031-5.25l1.032-1.786l-1.733-1l5.531-2.58ZM13.75 2.97a3.5 3.5 0 0 1 1.28 1.28l1.031 1.786l1.733-1l-.532 6.08l-5.53-2.58l1.73-1l-1.03-1.786a.5.5 0 0 0-.814-.074l-.052.074l-1.53 2.652l-2.598-1.5l1.53-2.652a3.5 3.5 0 0 1 4.782-1.28Z"/>`),
		g.Group(children),
	)
}