package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryCardOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMemoryCardOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M6 9a3 3 0 0 1 3-3h21.336a3 3 0 0 1 2.122.879l3.858 3.858l4.805 4.805A3 3 0 0 1 42 17.664V39a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Z"/><path fill="#555" d="M31 26H17a3 3 0 0 0-3 3v13h20V29a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" d="M29 16H17a3 3 0 0 1-3-3V6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMemoryCardOne0)"/>`),
		g.Group(children),
	)
}