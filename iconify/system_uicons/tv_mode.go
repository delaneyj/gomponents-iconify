package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.493 5.534l10-.036a2 2 0 0 1 2.007 2V12.5a2 2 0 0 1-2 2h-10a2 2 0 0 1-2-2V7.534a2 2 0 0 1 1.993-2z"/><path fill="currentColor" fill-rule="nonzero" d="M12.467 9.6L9.8 7.6A.5.5 0 0 0 9 8v4a.5.5 0 0 0 .8.4l2.667-2a.5.5 0 0 0 0-.8z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.464 16.5H15.5"/></g>`),
		g.Group(children),
	)
}