package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusWifiLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3v2H5v7h16v8h-1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1H7v1a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1H3v-8H2V8h1V5a2 2 0 0 1 2-2h7Zm7 11H5v4h14v-4Zm-9 1v2H6v-2h4Zm8 0v2h-4v-2h4Zm.5-14a4.5 4.5 0 1 1 0 9a4.5 4.5 0 0 1 0-9Zm0 5.167c-.491 0-.94.177-1.289.47l-.125.115L18.5 8.167l1.413-1.415a1.994 1.994 0 0 0-1.413-.585Zm0-2.667a4.65 4.65 0 0 0-3.128 1.203l-.173.165l.944.942a3.323 3.323 0 0 1 2.357-.977a3.32 3.32 0 0 1 2.201.83l.156.147l.943-.943A4.652 4.652 0 0 0 18.5 3.5Z"/>`),
		g.Group(children),
	)
}