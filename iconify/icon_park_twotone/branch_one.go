package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBranchOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 33V15"/><path fill="#555" d="M10 9h28v6H10z"/><path d="m8 32l6-7h19.974L40 32"/><path fill="#555" d="M4 33h8v8H4zm16 0h8v8h-8zm16 0h8v8h-8z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBranchOne0)"/>`),
		g.Group(children),
	)
}