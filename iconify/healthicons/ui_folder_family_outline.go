package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiFolderFamilyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M18.5 24a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm14 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM31 29a3 3 0 0 0-3 3v2h9v-2a3 3 0 0 0-3-3h-3Zm-17 0a3 3 0 0 0-3 3v2h9v-2a3 3 0 0 0-3-3h-3Zm12-2.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM24 30a3 3 0 0 0-3 3v1h6v-1a3 3 0 0 0-3-3Z"/><path fill-rule="evenodd" d="M24.712 10.555L27.154 15H39a3 3 0 0 1 3 3v18a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V12a3 3 0 0 1 3-3h13.082a3 3 0 0 1 2.63 1.555ZM9 37a1 1 0 0 1-1-1V17h31a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H9Zm15.872-22l-1.913-3.482a1 1 0 0 0-.877-.518H9a1 1 0 0 0-1 1v3h16.872Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}