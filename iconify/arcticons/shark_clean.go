package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharkClean(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.325 18.779l-1.488 8.442m.613-3.482a2.627 2.627 0 0 1 2.484-2.11h0a1.722 1.722 0 0 1 1.738 2.11l-.614 3.482m9.056-3.482a2.627 2.627 0 0 1 2.483-2.11h0m-2.111-.001l-.986 5.593m5.68-8.442l-1.489 8.442m.317-1.793l4.492-3.803m-3.062 2.593l2.476 2.991M9.5 26.296a1.962 1.962 0 0 0 1.907.925h1.25a2.624 2.624 0 0 0 2.478-2.11h0A1.72 1.72 0 0 0 13.401 23H12.02a1.72 1.72 0 0 1-1.734-2.11h0a2.624 2.624 0 0 1 2.478-2.111h1.25a1.962 1.962 0 0 1 1.907.925m11.841 5.406a2.627 2.627 0 0 1-2.483 2.111h0a1.722 1.722 0 0 1-1.738-2.11l.242-1.372a2.627 2.627 0 0 1 2.482-2.11h0a1.722 1.722 0 0 1 1.739 2.11m-.614 3.482l.986-5.593"/>`),
		g.Group(children),
	)
}