package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileUserOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 9a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 8.5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1Z"/><path fill="currentColor" fill-rule="evenodd" d="M7 2.25A2.75 2.75 0 0 0 4.25 5v14A2.75 2.75 0 0 0 7 21.75h10A2.75 2.75 0 0 0 19.75 19V8.198a1.75 1.75 0 0 0-.328-1.02L16.408 2.98a1.75 1.75 0 0 0-1.421-.73H7ZM5.75 5c0-.69.56-1.25 1.25-1.25h7.25v4.397c0 .414.336.75.75.75h3.25V19c0 .69-.56 1.25-1.25 1.25H7c-.69 0-1.25-.56-1.25-1.25V5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}