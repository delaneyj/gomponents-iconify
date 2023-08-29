package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStraightCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilFlagStraightCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M19.58 6H6.5v8h13.08l-2.299-3.283a1.25 1.25 0 0 1 0-1.434L19.58 6ZM6.5 5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h13.08a1 1 0 0 0 .819-1.573L18.1 10.143a.25.25 0 0 1 0-.286l2.3-3.284A1 1 0 0 0 19.579 5H6.5Z"/><path d="M6 6a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-1 0v-14A.5.5 0 0 1 6 6Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilFlagStraightCircleFilled0)"/></g>`),
		g.Group(children),
	)
}