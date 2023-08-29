package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockKeyholeMinimalisticUnlockedBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2.75A5.25 5.25 0 0 0 6.75 8v2.004C7.133 10 7.548 10 8 10h8c2.828 0 4.243 0 5.121.879C22 11.757 22 13.172 22 16c0 2.828 0 4.243-.879 5.121C20.243 22 18.828 22 16 22H8c-2.828 0-4.243 0-5.121-.879C2 20.243 2 18.828 2 16c0-2.828 0-4.243.879-5.121c.53-.531 1.256-.741 2.371-.824V8a6.75 6.75 0 0 1 13.287-1.687a.75.75 0 1 1-1.452.374A5.252 5.252 0 0 0 12 2.75ZM12.75 14a.75.75 0 0 0-1.5 0v4a.75.75 0 0 0 1.5 0v-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}