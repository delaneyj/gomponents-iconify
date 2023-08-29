package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Package(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.371.571L12 .423l-.371.148l-10 4L1 4.823v14.354l.629.251l10 4l.371.149l.371-.149l10-4l.629-.251V4.823l-.629-.252l-10-4ZM3 6.977v10.846l8 3.2V10.177l-8-3.2Zm10 3.2v10.846l8-3.2V6.977l-8 3.2ZM19.307 5.5L12 2.577L9.943 3.4l7.307 2.923l2.057-.823Zm-14.614 0L7.25 4.477L14.557 7.4L12 8.423L4.693 5.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}