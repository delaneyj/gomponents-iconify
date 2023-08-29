package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m251.2 118.4l-48-36A12 12 0 0 0 184 92v24H76V76h26.06a36 36 0 1 0 0-24H72a20 20 0 0 0-20 20v44H12a12 12 0 0 0 0 24h40v44a20 20 0 0 0 20 20h28v4a20 20 0 0 0 20 20h32a20 20 0 0 0 20-20v-32a20 20 0 0 0-20-20h-32a20 20 0 0 0-20 20v4H76v-40h108v24a12 12 0 0 0 19.2 9.6l48-36a12 12 0 0 0 0-19.2ZM136 52a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm-12 128h24v24h-24Zm84-40v-24l16 12Z"/>`),
		g.Group(children),
	)
}