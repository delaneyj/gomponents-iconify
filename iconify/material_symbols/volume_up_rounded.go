package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 11.975q0-2.075-1.1-3.787t-2.95-2.563q-.375-.175-.55-.537t-.05-.738q.15-.4.537-.575t.788 0Q18.1 4.85 19.55 7.063T21 11.974q0 2.7-1.45 4.913t-3.875 3.287q-.4.175-.788 0t-.537-.575q-.125-.375.05-.738t.55-.537q1.85-.85 2.95-2.563t1.1-3.787ZM7 15H4q-.425 0-.713-.288T3 14v-4q0-.425.288-.713T4 9h3l3.3-3.3q.475-.475 1.088-.213t.612.938v11.15q0 .675-.613.938T10.3 18.3L7 15Zm9.5-3q0 1.05-.475 1.988t-1.25 1.537q-.25.15-.513.013T14 15.1V8.85q0-.3.263-.438t.512.013q.775.625 1.25 1.575t.475 2Z"/>`),
		g.Group(children),
	)
}