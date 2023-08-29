package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm1082 299H374q-31 0-53 22t-22 53v149q0 31 22 52.5t53 21.5h373v821q0 31 22 53t53 22h149q31 0 52.5-22t21.5-53V597h373q31 0 53-21.5t22-52.5V374q0-31-22-53t-53-22z"/>`),
		g.Group(children),
	)
}