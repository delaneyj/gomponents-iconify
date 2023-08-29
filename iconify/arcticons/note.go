package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Note(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.348 11.35l-3.6-3.6h-20.2v35.8h29.6v-18.7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.948 10.55l-13 18.2l-1.5 12.6l11.5-5.3l13.3-18.4a.978.978 0 0 0-.2-1.4l-8.4-6a1.268 1.268 0 0 0-1.7.3Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.948 36.85a4.4 4.4 0 0 1 3.5 2.6m-2.5-10.7l10 7.3m1.2-22.9l10 7.4m-11.4-5.6l10.1 7.4m-26.6-1.8v-9.5c0-1.4 2.5-1.3 2.5 0v10.7c0 2.5-4.7 2.5-4.7 0V7.45c0-4 5.6-4 5.9 0"/>`),
		g.Group(children),
	)
}