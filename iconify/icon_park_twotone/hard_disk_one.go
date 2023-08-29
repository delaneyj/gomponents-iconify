package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDiskOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHardDiskOne0"><g fill="none" stroke="#fff" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M32 6H16v10h16V6Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 36h14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHardDiskOne0)"/>`),
		g.Group(children),
	)
}