package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1493 1419V373q0-31-22-52.5t-53-21.5h-149q-31 0-52.5 21.5T1195 373v374H597V373q0-31-21.5-52.5T523 299H374q-31 0-53 21.5T299 373v1046q0 31 22 52.5t53 21.5h149q31 0 52.5-21.5T597 1419v-374h598v374q0 31 21.5 52.5t52.5 21.5h149q31 0 53-21.5t22-52.5zm299-1083v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0h1120q139 0 237.5 98.5T1792 336z"/>`),
		g.Group(children),
	)
}