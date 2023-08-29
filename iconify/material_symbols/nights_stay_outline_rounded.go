package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightsStayOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.05 22q-.825 0-1.588-.113T8.9 21.5q.525-.275.938-.7t.712-.95q.35.075.763.113T12.1 20q1.325 0 2.6-.425t2.15-1.225q-2.15-.725-3.825-2.113t-2.737-3.225Q9.225 11.175 8.85 9.025T8.9 4.65q-2.225.8-3.563 2.938T4 12q-.525 0-1.038.113T2 12.45q-.075-2 .588-3.8t1.887-3.2q1.225-1.4 2.9-2.3T11 2.05q.4-.05.563.2t-.013.65q-1.05 2.325-.863 4.725t1.326 4.388Q13.15 14 15.124 15.35t4.525 1.6q.425.05.563.313t-.113.587q-1.425 2-3.538 3.075T12.05 22ZM7 20H4q-1.25 0-2.125-.863T1 17q0-1.25.875-2.125T4 14q.975 0 1.738.563T6.8 16H7q.825 0 1.413.588T9 18q0 .825-.588 1.413T7 20Zm3.425-7.675Z"/>`),
		g.Group(children),
	)
}