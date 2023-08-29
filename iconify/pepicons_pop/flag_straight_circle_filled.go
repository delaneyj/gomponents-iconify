package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStraightCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFlagStraightCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M7.5 7v6h10.159l-1.197-1.71a2.25 2.25 0 0 1 0-2.58L17.659 7H7.5Zm-1-2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h13.08a1 1 0 0 0 .819-1.573L18.1 10.143a.25.25 0 0 1 0-.286l2.3-3.284A1 1 0 0 0 19.579 5H6.5Z"/><path d="M6.5 6a1 1 0 0 1 1 1v13a1 1 0 1 1-2 0V7a1 1 0 0 1 1-1Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFlagStraightCircleFilled0)"/></g>`),
		g.Group(children),
	)
}