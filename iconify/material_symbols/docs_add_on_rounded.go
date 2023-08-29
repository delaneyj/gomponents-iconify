package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocsAddOnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20.975q-.425 0-.713-.288T16 19.976v-2h-2q-.425 0-.713-.288T13 16.976q0-.425.288-.713t.712-.287h2v-2q0-.425.288-.712t.712-.288q.425 0 .713.288t.287.712v2h2q.425 0 .713.288t.287.712q0 .425-.288.713t-.712.287h-2v2q0 .425-.288.713t-.712.287ZM5 18q-.425 0-.713-.288T4 17q0-.425.288-.713T5 16h6.075q-.075.525-.063 1t.088 1H5Zm0-4q-.425 0-.713-.288T4 13q0-.425.288-.713T5 12h8.65q-.575.4-1.038.9T11.8 14H5Zm0-4q-.425 0-.713-.288T4 9q0-.425.288-.713T5 8h13q.425 0 .713.288T19 9q0 .425-.288.713T18 10H5Zm0-4q-.425 0-.713-.288T4 5q0-.425.288-.713T5 4h13q.425 0 .713.288T19 5q0 .425-.288.713T18 6H5Z"/>`),
		g.Group(children),
	)
}