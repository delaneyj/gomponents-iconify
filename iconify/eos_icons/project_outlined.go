package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M6 17h12v2H6zm4-5.17l2.792 2.794l3.932-3.935L18 12V8h-4l1.31 1.275l-2.519 2.519L10 9l-4 4l1.414 1.414L10 11.83z"/><path fill="currentColor" d="M19 3h-3.298a4.497 4.497 0 0 0-.32-.425l-.01-.012a4.426 4.426 0 0 0-2.89-1.518a2.577 2.577 0 0 0-.964 0a4.426 4.426 0 0 0-2.89 1.518l-.01.012a4.497 4.497 0 0 0-.32.424V3H5a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V6a3.003 3.003 0 0 0-3-3Zm1 17a1 1 0 0 1-1 1H5a1.001 1.001 0 0 1-1-1V6a1.001 1.001 0 0 1 1-1h4.55a2.5 2.5 0 0 1 4.9 0H19a1.001 1.001 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}