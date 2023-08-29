package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHiding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 44h28a2 2 0 0 0 2-2V14H30V4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2ZM30 4l10 10"/><path d="M16 23a8.64 8.64 0 0 0 1.255 2.517C18.783 27.63 21.235 29 24 29s5.217-1.37 6.745-3.483A8.64 8.64 0 0 0 32 23m-10.479 6.068l-1.035 3.864m6-3.864l1.036 3.864m2.832-5.578l2.828 2.828M15 30.01l2.828-2.828"/></g>`),
		g.Group(children),
	)
}