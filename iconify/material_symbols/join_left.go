package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19q-2.925 0-4.963-2.038T1 12q0-2.925 2.038-4.963T8 5q.675 0 1.313.125t1.237.35Q9.1 6.7 8.3 8.4T7.5 12q0 1.9.8 3.6t2.25 2.925q-.6.225-1.238.35T8 19Zm8 0q-.675 0-1.313-.125t-1.237-.35q.425-.35.8-.762t.7-.863q.275.05.525.075T16 17q2.075 0 3.538-1.463T21 12q0-2.075-1.463-3.538T16 7q-.275 0-.525.025t-.525.075q-.325-.45-.7-.862t-.8-.763q.6-.225 1.238-.35T16 5q2.925 0 4.963 2.038T23 12q0 2.925-2.038 4.963T16 19Zm-4-1.25q-1.425-.975-2.212-2.5T9 12q0-1.725.788-3.25T12 6.25q1.425.975 2.212 2.5T15 12q0 1.725-.788 3.25T12 17.75Z"/>`),
		g.Group(children),
	)
}