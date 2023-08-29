package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Factory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#BF360C" d="M41.2 5h-7.3L32 43h11L41.2 5z"/><path fill="#E64A19" d="M33 23h-4v-6l-12 6v-6L5 23v20h28V23z"/><path fill="#FFC107" d="M9 27h4v4H9zm8 0h4v4h-4zm8 0h4v4h-4zM9 35h4v4H9zm8 0h4v4h-4zm8 0h4v4h-4z"/>`),
		g.Group(children),
	)
}