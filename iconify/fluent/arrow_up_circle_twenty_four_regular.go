package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpCircleTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M7.47 12.28l.084.074a.75.75 0 0 0 .977-.073l2.72-2.72v6.69l.006.102a.75.75 0 0 0 .743.648l.102-.007a.75.75 0 0 0 .648-.743V9.56l2.72 2.721l.085.073a.75.75 0 0 0 .977-1.133l-4-4.002l-.085-.072a.75.75 0 0 0-.976.072L7.47 11.22l-.073.084a.75.75 0 0 0 .073.977z" fill="currentColor"/><path d="M22 12c0-5.522-4.476-10-10-10C6.479 2 2 6.479 2 12c0 5.524 4.478 10 10 10c5.524 0 10-4.476 10-10zM3.5 12a8.5 8.5 0 1 1 17 0a8.5 8.5 0 0 1-17 0z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}