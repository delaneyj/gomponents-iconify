package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardDownloadOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 13.15V10q0-.425-.288-.713T12 9q-.425 0-.713.288T11 10v3.15l-.9-.875Q9.825 12 9.412 12t-.712.3q-.275.275-.275.7t.275.7l2.6 2.6q.3.3.7.3t.7-.3l2.6-2.6q.275-.275.287-.687T15.3 12.3q-.275-.275-.688-.288t-.712.263l-.9.875ZM6 22q-.825 0-1.413-.588T4 20V8.825q0-.4.15-.762t.425-.638l4.85-4.85q.275-.275.638-.425t.762-.15H18q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H6Zm0-2h12V4h-7.15L6 8.85V20Zm0 0h12H6Z"/>`),
		g.Group(children),
	)
}