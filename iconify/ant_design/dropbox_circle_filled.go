package ant_design

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropboxCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M663.8 455.5zm-151.5-93.8l-151.8 93.8l151.8 93.9l151.5-93.9zM512 64C264.6 64 64 264.6 64 512s200.6 448 448 448s448-200.6 448-448S759.4 64 512 64zm151.2 595.5L512.6 750l-151-90.5v-33.1l45.4 29.4l105.6-87.7l105.6 87.7l45.1-29.4v33.1zm-45.6-22.4l-105.3-87.7L407 637.1l-151-99.2l104.5-82.4L256 371.2L407 274l105.3 87.7L617.6 274L768 372.1l-104.2 83.5L768 539l-150.4 98.1z"/>`),
		g.Group(children),
	)
}