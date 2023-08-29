package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModifiedDateOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.372 11.013l.614.614l-6.04 6.04h-.613v-.614l6.04-6.04M15.771 7a.667.667 0 0 0-.467.193l-1.22 1.22l2.5 2.5l1.22-1.22a.664.664 0 0 0 0-.94l-1.56-1.56A.655.655 0 0 0 15.772 7Zm-2.4 2.127L6 16.5V19h2.5l7.372-7.373Z"/><path fill="currentColor" d="M19 1h-2v2H7V1H5v2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-1V1ZM4 21V5h16v16Z"/>`),
		g.Group(children),
	)
}