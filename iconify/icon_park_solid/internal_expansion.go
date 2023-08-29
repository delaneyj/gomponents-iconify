package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InternalExpansion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSInternalExpansion0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 42h32a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v32a2 2 0 0 0 2 2Z"/><path fill="#fff" fill-rule="evenodd" d="M42 8a2 2 0 0 0-2-2H28v14h14V8Z" clip-rule="evenodd"/><path d="m13 35l10-10m0 0v7m0-7h-7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSInternalExpansion0)"/>`),
		g.Group(children),
	)
}