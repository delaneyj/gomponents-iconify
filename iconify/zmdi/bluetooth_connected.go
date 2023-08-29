package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothConnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m85 216l-42 43l-43-43l43-43zm229-92l-92 92l92 92l-122 121h-21V267l-98 98l-30-30l119-119L43 97l30-30l98 98V3h21zM213 84v81l40-41zm40 224l-40-41v81zm88-135l43 43l-43 43l-42-43z"/>`),
		g.Group(children),
	)
}