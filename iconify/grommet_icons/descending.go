package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Descending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m11.92 16.714l-.354.353l-.354-.353L7 12.502l.707-.708l3.359 3.359V7h1v8.153l3.358-3.359l.707.708l-4.212 4.212Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}