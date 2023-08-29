package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boxes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 4v11h12V4H10zm2 2h2v4l2-2l2 2V6h2v7h-8V6zM3 17v11h12V17H3zm14 0v11h12V17H17zM5 19h2v4l2-2l2 2v-4h2v7H5v-7zm14 0h2v4l2-2l2 2v-4h2v7h-8v-7z"/>`),
		g.Group(children),
	)
}