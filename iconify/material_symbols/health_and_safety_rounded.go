package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealthAndSafetyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 13v1.5q0 .425.288.713t.712.287h1q.425 0 .713-.288t.287-.712V13H15q.425 0 .713-.288T16 12v-1q0-.425-.288-.713T15 10h-1.5V8.5q0-.425-.288-.713T12.5 7.5h-1q-.425 0-.713.288T10.5 8.5V10H9q-.425 0-.713.288T8 11v1q0 .425.288.713T9 13h1.5Zm1.5 8.9q-.175 0-.325-.025t-.3-.075Q8 20.675 6 17.637T4 11.1V6.375q0-.625.363-1.125t.937-.725l6-2.25q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375V11.1q0 3.5-2 6.538T12.625 21.8q-.15.05-.3.075T12 21.9Z"/>`),
		g.Group(children),
	)
}