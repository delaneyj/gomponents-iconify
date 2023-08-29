package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualSpace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><circle cx="11.999" cy="7.493" r="1.493" fill="currentColor"/><circle cx="11.999" cy="16.493" r="1.493" fill="currentColor"/><circle cx="16.496" cy="12" r="1.493" fill="currentColor"/><circle cx="7.496" cy="12" r="1.493" fill="currentColor"/><circle cx="11.999" cy="12" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}