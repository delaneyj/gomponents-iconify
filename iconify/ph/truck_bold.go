package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m251.14 115.54l-14-35A19.89 19.89 0 0 0 218.58 68H188v-4a12 12 0 0 0-12-12H24A20 20 0 0 0 4 72v112a20 20 0 0 0 20 20h14.06a36 36 0 0 0 67.88 0h44.12a36 36 0 0 0 67.88 0H232a20 20 0 0 0 20-20v-64a21.7 21.7 0 0 0-.86-4.46ZM188 92h27.88l6.4 16H188ZM72 204a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm92-41.92A36.32 36.32 0 0 0 150.06 180h-44.12a36 36 0 0 0-67.88 0H28v-40h136Zm0-46.08H28V76h136Zm20 88a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm44-24h-10.06A36.09 36.09 0 0 0 188 156.23V132h40Z"/>`),
		g.Group(children),
	)
}