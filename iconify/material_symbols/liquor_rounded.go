package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiquorRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.425 0-.713-.288T3 21q0-.425.288-.713T4 20h1v-3.2q-.875-.3-1.438-1.063T3 14V7q0-.425.288-.713T4 6h4q.425 0 .713.288T9 7v7q0 .975-.563 1.738T7 16.8V20h1q.425 0 .713.288T9 21q0 .425-.288.713T8 22H4Zm1-11h2V8H5v3Zm8 11q-.825 0-1.413-.588T11 20v-9.55q0-.65.375-1.162t.975-.738l.95-.35q.35-.125.525-.362T14 7.25V3q0-.425.288-.713T15 2h3q.425 0 .713.288T19 3v4.25q0 .35.175.588t.525.362l.95.35q.6.225.975.738T22 10.45V20q0 .825-.588 1.413T20 22h-7Zm3-17h1V4h-1v1Zm-3 7h7v-1.55l-.95-.35q-.95-.35-1.5-1.1T17 7.3V7h-1v.3q0 .95-.55 1.7t-1.5 1.1l-.95.35V12Zm0 8h7v-2h-7v2Z"/>`),
		g.Group(children),
	)
}