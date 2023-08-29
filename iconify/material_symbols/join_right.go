package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 19q-.675 0-1.313-.125t-1.237-.35Q14.9 17.3 15.7 15.6t.8-3.6q0-1.9-.8-3.6t-2.25-2.925q.6-.225 1.238-.35T16 5q2.925 0 4.963 2.038T23 12q0 2.925-2.038 4.963T16 19Zm-8 0q-2.925 0-4.963-2.038T1 12q0-2.925 2.038-4.963T8 5q.675 0 1.313.125t1.237.35q-.425.35-.8.763t-.7.862q-.25-.05-.513-.075T8 7Q5.925 7 4.463 8.463T3 12q0 2.075 1.463 3.538T8 17q.275 0 .537-.025t.513-.075q.325.45.7.863t.8.762q-.6.225-1.238.35T8 19Zm4-1.25q-1.425-.975-2.212-2.5T9 12q0-1.725.788-3.25T12 6.25q1.425.975 2.212 2.5T15 12q0 1.725-.788 3.25T12 17.75Z"/>`),
		g.Group(children),
	)
}