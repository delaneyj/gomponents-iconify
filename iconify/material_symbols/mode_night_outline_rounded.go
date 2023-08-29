package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeNightOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 4h-.525q-.25 0-.475.05q1.425 1.65 2.212 3.688T11.5 12q0 2.225-.788 4.263T8.5 19.95q.225.05.475.05H9.5q3.325 0 5.663-2.337T17.5 12q0-3.325-2.337-5.663T9.5 4Zm10 8q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T9.5 22q-.8 0-1.675-.15t-1.675-.425q-.3-.125-.475-.388T5.5 20.45q0-.225.088-.425t.262-.325q1.775-1.475 2.713-3.487T9.5 12q0-2.2-.938-4.213T5.85 4.3q-.175-.125-.262-.325T5.5 3.55q0-.325.175-.588t.475-.387q.8-.275 1.675-.425T9.5 2q2.075 0 3.9.787t3.175 2.138q1.35 1.35 2.138 3.175T19.5 12Zm-8 0Z"/>`),
		g.Group(children),
	)
}