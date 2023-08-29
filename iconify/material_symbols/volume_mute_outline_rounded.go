package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMuteOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15H8q-.425 0-.713-.288T7 14v-4q0-.425.288-.713T8 9h3l3.3-3.3q.475-.475 1.088-.213t.612.938v11.15q0 .675-.613.938T14.3 18.3L11 15Zm-2-2h2.85L14 15.15v-6.3L11.85 11H9v2Zm2.5-1Z"/>`),
		g.Group(children),
	)
}