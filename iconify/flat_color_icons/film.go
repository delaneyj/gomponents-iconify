package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Film(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="M45 9H3v30h42V9zM22 37v-4h4v4h-4zm8 0v-4h4v4h-4zm8 0v-4h4v4h-4zm-24 0v-4h4v4h-4zm-8 0v-4h4v4H6zm16-22v-4h4v4h-4zm8 0v-4h4v4h-4zm8 0v-4h4v4h-4zm-24 0v-4h4v4h-4zm-8 0v-4h4v4H6z"/>`),
		g.Group(children),
	)
}