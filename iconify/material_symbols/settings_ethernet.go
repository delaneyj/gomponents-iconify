package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsEthernet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 18l-1.4-1.4l4.55-4.6l-4.55-4.6L17 6l6 6l-6 6ZM7 18l-6-6l6-6l1.4 1.4L3.85 12l4.55 4.6L7 18Zm1-5q-.425 0-.713-.288T7 12q0-.425.288-.713T8 11q.425 0 .713.288T9 12q0 .425-.288.713T8 13Zm4 0q-.425 0-.713-.288T11 12q0-.425.288-.713T12 11q.425 0 .713.288T13 12q0 .425-.288.713T12 13Zm4 0q-.425 0-.713-.288T15 12q0-.425.288-.713T16 11q.425 0 .713.288T17 12q0 .425-.288.713T16 13Z"/>`),
		g.Group(children),
	)
}