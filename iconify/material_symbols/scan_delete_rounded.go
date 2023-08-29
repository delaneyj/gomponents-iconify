package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanDeleteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 9l-5-5v3q0 .825.588 1.413T15 9h3Zm0 10.425L16.6 20.8q-.275.275-.688.288T15.2 20.8q-.275-.275-.275-.7t.275-.7l1.4-1.4l-1.4-1.4q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.4 1.4l1.4-1.4q.275-.275.688-.287t.712.287q.275.275.275.7t-.275.7L19.425 18l1.375 1.4q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L18 19.425ZM6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762v3.525q-.475-.175-.988-.263T17.976 12q-2.5 0-4.237 1.738T12 17.974q0 1.125.4 2.163T13.55 22H6Z"/>`),
		g.Group(children),
	)
}