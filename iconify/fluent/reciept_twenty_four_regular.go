package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecieptTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M4 6.25A2.25 2.25 0 0 1 6.25 4h8.5A2.25 2.25 0 0 1 17 6.25V14h3.5v3.25a3.25 3.25 0 0 1-3.25 3.25h-10A3.25 3.25 0 0 1 4 17.25v-11zm13 9.25V19h.25A1.75 1.75 0 0 0 19 17.25V15.5h-2zM15.5 19V6.25a.75.75 0 0 0-.75-.75h-8.5a.75.75 0 0 0-.75.75v11c0 .966.784 1.75 1.75 1.75h8.25z" fill="currentColor"/><path d="M7 8.75A.75.75 0 0 1 7.75 8h5.5a.75.75 0 0 1 0 1.5h-5.5A.75.75 0 0 1 7 8.75zm0 3.5a.75.75 0 0 1 .75-.75h5.5a.75.75 0 0 1 0 1.5h-5.5a.75.75 0 0 1-.75-.75zm0 3.5a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}