package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 28H6a2.002 2.002 0 0 1-2-2V11a2.002 2.002 0 0 1 2-2h5.667a2.012 2.012 0 0 1 1.2.4l3.466 2.6H26a2.002 2.002 0 0 1 2 2v12a2.002 2.002 0 0 1-2 2zM11.666 11H5.998L6 26h20V14H15.666zM28 9H17.666l-4-3H6V4h7.667a2.012 2.012 0 0 1 1.2.4L18.333 7H28z"/>`),
		g.Group(children),
	)
}