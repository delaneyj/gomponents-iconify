package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Privacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.55 23.35L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4ZM22 17.5l-4-4v1.675l-6-6V9q0-.825-.588-1.413T10 7q-.05 0-.075.013t-.075.012L6.825 4H16q.825 0 1.413.588T18 6v4.5l4-4v11ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4l14 14q0 .825-.588 1.413T16 20h-5.3v-3.05q.925-.125 1.725-.563t1.4-1.162L12.8 14.2q-.5.65-1.238 1.012T10 15.575q-1.5 0-2.538-1.037T6.426 12H5q0 1.875 1.213 3.275T9.3 16.95V20H4Zm4-9.2V12q0 .825.588 1.413T10 14q.25 0 .475-.075t.45-.2L8 10.8Z"/>`),
		g.Group(children),
	)
}