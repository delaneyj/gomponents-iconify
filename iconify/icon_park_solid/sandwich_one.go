package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SandwichOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSandwichOne0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M17.799 40.142s4.16 4.16 8.652 2.995c4.492-1.165 15.805-12.479 16.97-16.97c1.165-4.493-2.995-8.652-2.995-8.652M7.9 30.243s-4.16-4.16-2.995-8.652c1.164-4.492 12.478-15.806 16.97-16.97c4.492-1.165 8.652 2.994 8.652 2.994"/><rect width="40" height="14" x="5.071" y="33.071" fill="#fff" stroke="#fff" rx="7" transform="rotate(-45 5.071 33.071)"/><path stroke="#000" d="m15.678 33.779l.563-1.689a5.996 5.996 0 0 1 3.326-3.615l.42-.18a4.98 4.98 0 0 0 2.762-3.002v0a4.98 4.98 0 0 1 2.762-3.002l.95-.407a4.712 4.712 0 0 0 2.614-2.84v0a4.712 4.712 0 0 1 2.98-2.981l2.008-.67"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSandwichOne0)"/>`),
		g.Group(children),
	)
}