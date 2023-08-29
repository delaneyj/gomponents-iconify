package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisabledPicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M44 23.999a2 2 0 0 0-4 0h4Zm-20-16a2 2 0 1 0 0-4v4Zm15 32H9v4h30v-4Zm-31-1v-30H4v30h4Zm32-15v15h4v-15h-4Zm-31-16h15v-4H9v4Zm0 32a1 1 0 0 1-1-1H4a5 5 0 0 0 5 5v-4Zm30 4a5 5 0 0 0 5-5h-4a1 1 0 0 1-1 1v4Zm-31-35a1 1 0 0 1 1-1v-4a5 5 0 0 0-5 5h4Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 35l10.693-9.802a2 2 0 0 1 2.653-.044L32 36m-4-5l4.773-4.773a2 2 0 0 1 2.615-.186L42 31"/><circle cx="36" cy="12" r="6" stroke="currentColor" stroke-width="4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m32 8l8 8"/></g>`),
		g.Group(children),
	)
}