package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackupOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 20q-2.275 0-3.888-1.575T1 14.575q0-1.95 1.175-3.475T5.25 9.15q.625-2.3 2.5-3.725T12 4q2.925 0 4.963 2.038T19 11q1.725.2 2.863 1.488T23 15.5q0 1.875-1.313 3.188T18.5 20H13q-.825 0-1.413-.588T11 18v-5.2l-.9.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l2.6-2.6q.15-.15.325-.212T12 9.425q.2 0 .375.063t.325.212l2.6 2.6q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-.9-.9V18h5.5q1.05 0 1.775-.725T21 15.5q0-1.05-.725-1.775T18.5 13H17v-2q0-2.075-1.463-3.538T12 6Q9.925 6 8.462 7.463T7 11h-.5q-1.45 0-2.475 1.025T3 14.5q0 1.45 1.025 2.475T6.5 18H8q.425 0 .713.288T9 19q0 .425-.288.713T8 20H6.5Zm5.5-7Z"/>`),
		g.Group(children),
	)
}