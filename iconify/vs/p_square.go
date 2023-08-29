package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm339 1054h334q147 0 251-110t104-267q0-154-102-266q-104-112-253-112H502q-30 0-52 22t-22 53v1044q0 31 22 53t52 22h98q31 0 53-22t22-53v-364zm0-531h334q50 0 85 45t35 109q0 63-35.5 108t-84.5 45H675V523z"/>`),
		g.Group(children),
	)
}