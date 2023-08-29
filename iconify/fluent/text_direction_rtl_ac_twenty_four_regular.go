package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionRtlAcTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M18.25 4a.75.75 0 0 1 .75.75V8h1.25a.75.75 0 0 1 0 1.5H19v3.75a.75.75 0 0 1-1.5 0v-8.5a.75.75 0 0 1 .75-.75zM9.5 5.75a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 .747.682c.194 2.135-1.141 5.97-5.531 7.286a.75.75 0 0 1-.431-1.436c3.132-.94 4.298-3.398 4.46-5.032H10.25a.75.75 0 0 1-.75-.75zm-4.28.47a.75.75 0 0 1 1.06 1.06L5.56 8h2.19a.75.75 0 0 1 0 1.5H5.56l.72.72a.75.75 0 1 1-1.06 1.06l-2-2a.75.75 0 0 1 0-1.06l2-2zm0 8a.75.75 0 0 1 1.06 1.06l-.72.72h14.69a.75.75 0 0 1 0 1.5H5.56l.72.72a.75.75 0 1 1-1.06 1.06l-2-2a.75.75 0 0 1 0-1.06l2-2z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}