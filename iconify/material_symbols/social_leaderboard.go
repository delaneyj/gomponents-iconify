package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialLeaderboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.475 21Q5.2 21 3.6 19.4T2 15.525q0-2.15 1.438-3.713t3.587-1.737L3 2h7l2 4l2-4h7l-4 8.025q2.125.2 3.563 1.763T22 15.5q0 2.3-1.6 3.9T16.5 21q-.225 0-.463-.013t-.462-.062q1.375-.9 2.15-2.337T18.5 15.5q0-2.725-1.888-4.612T12 9q-2.725 0-4.612 1.888T5.5 15.5q0 1.7.7 3.2t2.2 2.225q-.225.05-.463.063T7.476 21ZM12 20q-1.875 0-3.188-1.313T7.5 15.5q0-1.875 1.313-3.188T12 11q1.875 0 3.188 1.313T16.5 15.5q0 1.875-1.313 3.188T12 20Zm-1.85-1.75l1.85-1.4l1.85 1.4l-.7-2.275L15 14.65h-2.275L12 12.25l-.725 2.4H9l1.85 1.325l-.7 2.275Z"/>`),
		g.Group(children),
	)
}