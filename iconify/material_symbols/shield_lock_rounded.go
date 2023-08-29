package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldLockRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 16h4q.425 0 .713-.288T15 15v-3q0-.425-.288-.713T14 11v-1q0-.825-.588-1.413T12 8q-.825 0-1.413.588T10 10v1q-.425 0-.713.288T9 12v3q0 .425.288.713T10 16Zm1-5v-1q0-.425.288-.713T12 9q.425 0 .713.288T13 10v1h-2Zm1 10.9q-.175 0-.325-.025t-.3-.075Q8 20.675 6 17.637T4 11.1V6.375q0-.625.363-1.125t.937-.725l6-2.25q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375V11.1q0 3.5-2 6.538T12.625 21.8q-.15.05-.3.075T12 21.9Z"/>`),
		g.Group(children),
	)
}