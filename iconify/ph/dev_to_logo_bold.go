package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevToLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 52H24A20 20 0 0 0 4 72v112a20 20 0 0 0 20 20h208a20 20 0 0 0 20-20V72a20 20 0 0 0-20-20Zm-4 128H28V76h200Zm-120-24v-56a12 12 0 0 1 12-12h20a12 12 0 0 1 0 24h-8v4a12 12 0 0 1 0 24v4h8a12 12 0 0 1 0 24h-20a12 12 0 0 1-12-12Zm52.46-52.7a12 12 0 1 1 23.08-6.6l4.46 15.62l4.46-15.62a12 12 0 0 1 23.08 6.6l-16 56a12 12 0 0 1-23.08 0ZM56 168h8a36 36 0 0 0 36-36v-8a36 36 0 0 0-36-36h-8a12 12 0 0 0-12 12v56a12 12 0 0 0 12 12Zm12-55.31A12 12 0 0 1 76 124v8a12 12 0 0 1-8 11.31Z"/>`),
		g.Group(children),
	)
}