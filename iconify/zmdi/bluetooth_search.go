package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m240 216l49-49q10 24 10 49q0 26-10 50zm113-113q31 51 31 111q0 61-33 113l-25-25q21-41 21-86q0-46-21-86zm-82 21l-92 92l92 92l-122 121h-21V267l-98 98l-30-30l119-119L0 97l30-30l98 98V3h21zM171 84v81l40-41zm40 224l-40-41v81z"/>`),
		g.Group(children),
	)
}