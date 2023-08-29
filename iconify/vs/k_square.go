package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm28 1419q0 31 22 52.5t53 21.5h149q31 0 53-21.5t22-52.5v-374h123q35 0 55.5 11t45.5 41l204 321q19 30 45 52.5t53 22.5h188q55 0 56-32q0-17-38-83l-314-482l314-482q38-66 38-83q-1-32-56-32h-188q-27 0-53 22.5t-45 52.5L887 695q-25 30-45.5 41T786 747H663V373q0-31-22-52.5T588 299H439q-31 0-53 21.5T364 373v1046z"/>`),
		g.Group(children),
	)
}