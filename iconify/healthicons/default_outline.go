package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DefaultOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10 10a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v14.055a8.948 8.948 0 0 1 2 .457V10a4 4 0 0 0-4-4H12a4 4 0 0 0-4 4v24a4 4 0 0 0 4 4h11.515a8.968 8.968 0 0 1-1.003-2H12a2 2 0 0 1-2-2V10Z"/><path d="M12 14a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H13a1 1 0 0 1-1-1Zm0 5a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H13a1 1 0 0 1-1-1Zm1 4a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2h-8Zm-1 6a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2h-5a1 1 0 0 1-1-1Zm19 11a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0 2a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z"/><path d="M34.997 29.028a1 1 0 0 0-1.414 0l-.79.79l1.413 1.415l.791-.791a1 1 0 0 0 0-1.414ZM33.5 31.94l-1.414-1.414l-5.069 5.069v1.414h1.414l5.07-5.069Z"/></g>`),
		g.Group(children),
	)
}