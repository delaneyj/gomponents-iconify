package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMemoryOne0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 28h38v14H5zM5 6h38v14H5z"/><rect width="4" height="4" x="11" y="11" fill="#fff" rx="2"/><rect width="4" height="4" x="11" y="33" fill="#fff" rx="2"/><rect width="4" height="4" x="19" y="11" fill="#fff" rx="2"/><rect width="4" height="4" x="19" y="33" fill="#fff" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 13h4m-4 22h4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMemoryOne0)"/>`),
		g.Group(children),
	)
}