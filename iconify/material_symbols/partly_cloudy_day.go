package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartlyCloudyDay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20q-1.65 0-2.825-1.175T2 16q0-1.65 1.175-2.825T6 12q1.2 0 2.212.65t1.463 1.775l.25.575h.6q1.05 0 1.763.725T13 17.5q0 1.05-.725 1.775T10.5 20H6Zm8.975-2.8q-.1-1.575-1.137-2.725t-2.613-1.425q-.775-1.35-2.087-2.138T6.25 10q.65-1.825 2.225-2.913T12 6q2.5 0 4.25 1.75T18 12q0 1.625-.8 3.013T14.975 17.2ZM11 5V1h2v4h-2Zm6.65 2.75l-1.4-1.4l2.8-2.85l1.425 1.425L17.65 7.75ZM19 13v-2h4v2h-4Zm.05 7.5l-2.8-2.85l1.4-1.4l2.85 2.8l-1.45 1.45ZM6.35 7.75L3.525 4.925L4.95 3.5l2.8 2.85l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}