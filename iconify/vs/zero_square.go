package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeroSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm560 1493q150 0 276.5-80t200.5-217.5t74-299.5t-74-299.5T1172.5 379T896 299t-276.5 80T419 596.5T345 896t74 299.5T619.5 1413t276.5 80zm-77-271q40 10 75 10q127 0 216.5-97t89.5-239q0-107-56-195zm154-652q-40-10-75-10q-128 1-217 97.5T592 896q0 107 56 195z"/>`),
		g.Group(children),
	)
}