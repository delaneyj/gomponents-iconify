package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownloadOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 20q-2.275 0-3.888-1.575T1 14.575q0-1.95 1.175-3.475T5.25 9.15q.425-1.8 2.125-3.425T11 4.1q.825 0 1.413.588T13 6.1v6.05l.9-.875q.275-.275.688-.275t.712.3q.275.275.275.7t-.275.7l-2.6 2.6q-.15.15-.325.213t-.375.062q-.2 0-.375-.062T11.3 15.3l-2.6-2.6q-.275-.275-.287-.687T8.7 11.3q.275-.275.688-.288t.712.263l.9.875V6.1q-1.9.35-2.95 1.838T7 11h-.5q-1.45 0-2.475 1.025T3 14.5q0 1.45 1.025 2.475T6.5 18h12q1.05 0 1.775-.725T21 15.5q0-1.05-.725-1.775T18.5 13H17v-2q0-1.2-.55-2.238T15 7V4.675q1.85.875 2.925 2.588T19 11q1.725.2 2.863 1.488T23 15.5q0 1.875-1.313 3.188T18.5 20h-12Zm5.5-8.95Z"/>`),
		g.Group(children),
	)
}