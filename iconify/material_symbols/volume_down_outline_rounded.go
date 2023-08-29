package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 15H6q-.425 0-.713-.288T5 14v-4q0-.425.288-.713T6 9h3l3.3-3.3q.475-.475 1.088-.213t.612.938v11.15q0 .675-.613.938T12.3 18.3L9 15Zm9.5-3q0 1.05-.475 1.988t-1.25 1.537q-.25.15-.512.013T16 15.1V8.85q0-.3.263-.438t.512.013q.775.625 1.25 1.575t.475 2ZM12 8.85L9.85 11H7v2h2.85L12 15.15v-6.3ZM9.5 12Z"/>`),
		g.Group(children),
	)
}