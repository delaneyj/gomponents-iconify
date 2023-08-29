package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RespiratoryRateOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.35 17.6L16 15.925V15v.925L14.35 17.6Zm5.3 0L18 15.925V15v.925l1.65 1.675ZM12 12ZM2 9V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v5.5h-2V6H4v3H2Zm2 11q-.825 0-1.413-.588T2 18v-3h2v3h5.5v2H4Zm3-3q-.275 0-.525-.138T6.1 16.45L4.375 13H2v-2h3q.275 0 .525.138t.375.412l1.1 2.2l3.1-6.2q.125-.25.375-.375T11 7.05q.275 0 .525.125t.375.375l1.975 3.975q-.525.05-.988.225t-.862.5l-1.025-2l-3.1 6.2q-.125.275-.375.413T7 17Zm5.5 6q-.625 0-1.063-.438T11 21.5v-3.675l1.325-3.525q.225-.575.738-.938T14.2 13H16v-2h2v2h1.8q.625 0 1.137.363t.738.937L23 17.825V21.5q0 .625-.438 1.063T21.5 23h-2q-.625 0-1.063-.438T18 21.5V20h2v1h1v-2.825L19.8 15H18v.925l1.65 1.675l-1.4 1.4L17 17.75L15.75 19l-1.4-1.4L16 15.925V15h-1.8L13 18.175V21h1v-1h2v1.5q0 .625-.438 1.063T14.5 23h-2Zm4.5-5Z"/>`),
		g.Group(children),
	)
}