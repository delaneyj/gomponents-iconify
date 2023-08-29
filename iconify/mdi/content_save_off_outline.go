package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentSaveOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.2 5l-2-2H17l4 4v10.8l-2-2V7.83L16.17 5H8.2m6.8 5V6H9.2l4 4H15m7.11 11.46l-1.27 1.27L19.1 21H5a2 2 0 0 1-2-2V4.9L1.11 3l1.28-1.27l19.72 19.73m-5-2.46l-2.52-2.5a2.996 2.996 0 1 1-4.09-4.09L8.11 10H6V7.89l-1-1V19h12.11Z"/>`),
		g.Group(children),
	)
}