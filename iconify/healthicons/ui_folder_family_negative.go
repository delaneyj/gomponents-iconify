package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiFolderFamilyNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsUiFolderFamilyNegative0)"><path d="m24.872 15l-1.913-3.482a1 1 0 0 0-.877-.518H9a1 1 0 0 0-1 1v3h16.872ZM12.5 25a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm20 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM28 33a3 3 0 0 1 3-3h3a3 3 0 0 1 3 3v2h-9v-2Zm-17 0a3 3 0 0 1 3-3h3a3 3 0 0 1 3 3v2h-9v-2Zm13-3.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM21 34a3 3 0 1 1 6 0v1h-6v-1Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0Zm-9 15a3 3 0 0 1 3 3v18a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V12a3 3 0 0 1 3-3h13.082a3 3 0 0 1 2.63 1.555L27.154 15H39Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsUiFolderFamilyNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}