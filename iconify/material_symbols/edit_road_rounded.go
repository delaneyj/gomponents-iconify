package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditRoadRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11.9V5q0-.425.288-.713T17 4q.425 0 .713.288T18 5v4.9l-2 2ZM4 19V5q0-.425.288-.713T5 4q.425 0 .713.288T6 5v14q0 .425-.288.713T5 20q-.425 0-.713-.288T4 19Zm6-12V5q0-.425.288-.713T11 4q.425 0 .713.288T12 5v2q0 .425-.288.713T11 8q-.425 0-.713-.288T10 7Zm0 6v-2q0-.425.288-.713T11 10q.425 0 .713.288T12 11v2q0 .425-.288.713T11 14q-.425 0-.713-.288T10 13Zm0 6v-2q0-.425.288-.713T11 16q.425 0 .713.288T12 17v2q0 .425-.288.713T11 20q-.425 0-.713-.288T10 19Zm9.075-6.225l2.15 2.1l-4.65 4.675q-.2.2-.487.325T15.5 20h-.75q-.325 0-.537-.213T14 19.25v-.75q0-.3.113-.588t.312-.487l4.65-4.65Zm2.875 1.4l-2.125-2.125l.7-.7q.3-.3.725-.3t.7.3l.7.725q.275.3.275.713t-.275.687l-.7.7Z"/>`),
		g.Group(children),
	)
}