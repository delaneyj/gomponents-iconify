package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPowerCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M7.477 8.46a.6.6 0 1 1 .854.842a6.018 6.018 0 0 0-1.731 4.24c0 3.312 2.643 5.992 5.9 5.992c3.257 0 5.9-2.68 5.9-5.992a6.02 6.02 0 0 0-1.731-4.24a.6.6 0 1 1 .854-.842a7.218 7.218 0 0 1 2.077 5.082c0 3.97-3.177 7.192-7.1 7.192c-3.923 0-7.1-3.222-7.1-7.192c0-1.93.756-3.743 2.077-5.082Z"/><path d="M11.878 4.25a.6.6 0 0 1 1.2 0v7.085a.6.6 0 0 1-1.2 0V4.25Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPowerCircleFilled0)"/></g>`),
		g.Group(children),
	)
}