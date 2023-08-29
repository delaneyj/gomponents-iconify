package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11.445 2.168a1 1 0 0 1 1.325.194c.468.565 1.852 1.93 3.447 2.687A8.998 8.998 0 0 1 21 13A9 9 0 1 1 5.4 6.882a1 1 0 0 1 1.644.267c.127.282.471.663.992 1.066l.083.064c.2-1.41.632-2.592 1.145-3.538c.686-1.268 1.53-2.139 2.181-2.573zm.476 8.017a1 1 0 0 1 1.369.202c.042.054.158.186.323.334a2.536 2.536 0 0 0 .57.396c1.1.633 1.817 1.85 1.817 3.216C16 16.297 14.493 18 12.5 18S9 16.297 9 14.333c0-.943.341-1.812.912-2.469a1 1 0 0 1 .994-.314c.039-.092.08-.18.122-.264c.253-.498.582-.88.893-1.101z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}