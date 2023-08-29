package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 128h-78L275 896h45q26 0 45 18.5t19 45t-19 45.5t-45 19H64q-27 0-45.5-19T0 959.5t18.5-45T64 896h80l223-768h-47q-27 0-45.5-18.5t-18.5-45T274.5 19T320 0h256q26 0 45 19t19 45.5t-19 45t-45 18.5z"/>`),
		g.Group(children),
	)
}