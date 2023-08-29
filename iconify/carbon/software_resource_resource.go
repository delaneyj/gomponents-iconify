package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoftwareResourceResource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m6.899 10.283l3.394-3.394l1.414 1.414l-3.394 3.394zM4.9 19.7l-2.3-2.3C2.2 17 2 16.5 2 16s.2-1 .6-1.4l2.3-2.3l1.4 1.4L4 16l2.3 2.3l-1.4 1.4zm1.989 2.007l1.414-1.414l3.394 3.394l-1.414 1.414zM16 30c-.5 0-1-.2-1.4-.6l-2.3-2.3l1.4-1.4L16 28l2.3-2.3l1.4 1.4l-2.3 2.3c-.4.4-.9.6-1.4.6zm4.293-6.303l3.394-3.394l1.414 1.414l-3.394 3.394zM27.1 19.7l-1.4-1.4L28 16l-2.3-2.3l1.4-1.4l2.3 2.3c.4.4.6.9.6 1.4s-.2 1-.6 1.4l-2.3 2.3zM20.303 8.313l1.414-1.414l3.394 3.394l-1.414 1.414zM16 2c-.5 0-1 .2-1.4.6l-2.3 2.3l1.4 1.4L16 4l2.3 2.3l1.4-1.4l-2.3-2.3C17 2.2 16.5 2 16 2z"/>`),
		g.Group(children),
	)
}