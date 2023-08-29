package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m13 4l3-3l-1-1l-3 3H5V0H3v3H0v2h3v8h8v3h2v-3h3v-2h-3V4zM5 5h5l-5 5V5zm1 6l5-5v5H6z"/>`),
		g.Group(children),
	)
}