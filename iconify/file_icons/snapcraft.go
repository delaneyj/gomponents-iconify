package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snapcraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m294.5 122l112.9 50.2L294.5 285V122zM79 500.6l199.2-199.2l-60.8-60.4L79 500.6zM0 11.4l284.9 283.2V114.4L0 11.4zm465.1 103h-164L512 208.2l-46.9-93.8z"/>`),
		g.Group(children),
	)
}