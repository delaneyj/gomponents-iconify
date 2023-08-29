package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dianziqian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m33 8l-9 5l10 6v10.214L14 18v11l20 12l9-5.893V13.893L33 8Z"/><path d="m24 35l-9 6l-10-6V14l10-6l9 5"/></g>`),
		g.Group(children),
	)
}