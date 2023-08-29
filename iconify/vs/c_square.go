package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm-37 896q0 162 80 299.5T596.5 1413t299.5 80q192 0 340.5-98.5T1449 1127q30-81-23-81h-186q-20 0-32 4t-16 7.5t-15 19.5t-18 25q-47 60-116 95t-149 35q-140 0-238.5-98.5T557 896t98.5-237.5T894 560q80 0 149 35t116 95q6 8 13 18l10.5 15l8.5 10.5l11 7.5l15.5 3.5l22.5 1.5h186q53 0 23-81q-64-169-212.5-267.5T896 299q-162 0-299.5 80T379 596.5T299 896z"/>`),
		g.Group(children),
	)
}