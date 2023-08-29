package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 6v20h26V6H3zm2 2h4v4H5V8zm6 0h4v4h-4V8zm6 0h4v4h-4V8zm6 0h4v4h-4V8zM5 14h4v4H5v-4zm6 0h4v4h-4v-4zm6 0h4v4h-4v-4zm6 0h4v4h-4v-4zM5 20h4v4H5v-4zm6 0h4v4h-4v-4zm6 0h4v4h-4v-4zm6 0h4v4h-4v-4z"/>`),
		g.Group(children),
	)
}