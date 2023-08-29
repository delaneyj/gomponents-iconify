package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vagrant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M506.9 36.7L443.8 0L317.7 73.4v34.5l-64 166.8l-63.6-165V73.5L63.4 0L.1 36.6L0 73.9l158.6 383.6l95.1 54.4l94.9-54.7l158.7-384l-.4-36.5z"/>`),
		g.Group(children),
	)
}