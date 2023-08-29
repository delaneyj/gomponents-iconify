package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardholderFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 48H48a24 24 0 0 0-24 24v112a24 24 0 0 0 24 24h160a24 24 0 0 0 24-24V72a24 24 0 0 0-24-24Zm-56.48 76.81a24 24 0 0 1-47 0A16 16 0 0 0 88.81 112H40V96h176v16h-48.81a16 16 0 0 0-15.67 12.81ZM48 64h160a8 8 0 0 1 8 8v8H40v-8a8 8 0 0 1 8-8Z"/>`),
		g.Group(children),
	)
}