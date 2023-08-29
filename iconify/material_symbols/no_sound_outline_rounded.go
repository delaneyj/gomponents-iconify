package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSoundOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 13.4l-1.9 1.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l1.9-1.9l-1.9-1.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.9 1.9l1.9-1.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L19.4 12l1.9 1.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L18 13.4ZM7 15H4q-.425 0-.713-.288T3 14v-4q0-.425.288-.713T4 9h3l3.3-3.3q.475-.475 1.088-.213t.612.938v11.15q0 .675-.613.938T10.3 18.3L7 15Zm3-6.15L7.85 11H5v2h2.85L10 15.15v-6.3ZM7.5 12Z"/>`),
		g.Group(children),
	)
}