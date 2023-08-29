package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 128a71.69 71.69 0 0 1-15.78 44.91L83.09 71.78A71.95 71.95 0 0 1 200 128Zm-144 0a71.95 71.95 0 0 0 116.91 56.22L71.78 83.09A71.69 71.69 0 0 0 56 128Zm180 0A108 108 0 1 1 128 20a108.12 108.12 0 0 1 108 108Zm-20 0a88 88 0 1 0-88 88a88.1 88.1 0 0 0 88-88Z"/>`),
		g.Group(children),
	)
}