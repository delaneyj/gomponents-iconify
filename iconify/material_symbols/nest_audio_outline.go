package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestAudioOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21q-1.65 0-2.825-1.175T5 17V8q0-2.075 1.463-3.538T10 3h4q2.075 0 3.538 1.463T19 8v9q0 1.65-1.175 2.825T15 21H9Zm0-2h6q.825 0 1.413-.588T17 17V8q0-1.25-.875-2.125T14 5h-4q-1.25 0-2.125.875T7 8v9q0 .825.588 1.413T9 19Zm0-6q.425 0 .713-.288T10 12q0-.425-.288-.713T9 11q-.425 0-.713.288T8 12q0 .425.288.713T9 13Zm3 0q.425 0 .713-.288T13 12q0-.425-.288-.713T12 11q-.425 0-.713.288T11 12q0 .425.288.713T12 13Zm3 0q.425 0 .713-.288T16 12q0-.425-.288-.713T15 11q-.425 0-.713.288T14 12q0 .425.288.713T15 13Zm-3-1Z"/>`),
		g.Group(children),
	)
}