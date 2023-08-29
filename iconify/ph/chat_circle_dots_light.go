package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCircleDotsLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M138 128a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm-54-10a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm88 0a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm58 10a102 102 0 0 1-150.69 89.65l-34.87 11.62a14 14 0 0 1-17.71-17.71l11.62-34.87A102 102 0 1 1 230 128Zm-12 0a90 90 0 1 0-167.92 45.06a6 6 0 0 1 .5 4.91l-12.46 37.38a2 2 0 0 0 2.53 2.53L78 205.42a6.2 6.2 0 0 1 1.9-.31a6.09 6.09 0 0 1 3 .81A90 90 0 0 0 218 128Z"/>`),
		g.Group(children),
	)
}