package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireplaceOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.6 16.95q.975.8 2.113.388T14 15.65q.125-1.175-.725-1.738T12 12.45q-.125.35-.125.65t.075.65q.075.425.175.8t.025.8q-.125.45-.55.925t-1 .675ZM4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v16q0 .825-.588 1.413T20 22H4Zm8-4q1.25 0 2.125-.875T15 15q0-.6-.25-1t-.7-.75q-.95-.675-1.588-1.413T11.45 10.35q-1.1.875-1.775 1.988T9 14.95q0 .875.9 1.963T12 18Zm-8 2h2v-1q0-.425.288-.713T7 18h1.25q-.575-.725-.913-1.525T7 14.95q0-2.15.975-3.862T12.05 7.55q.375-.225.6-.087t.275.637q.075.9.625 1.813t1.6 1.687q.825.6 1.338 1.413T17 15q0 .875-.275 1.613T16 18h1q.425 0 .713.288T18 19v1h2V4H4v16Zm8-2Z"/>`),
		g.Group(children),
	)
}