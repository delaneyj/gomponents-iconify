package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm339 507v285h322q50 0 85-41.5t35-100.5t-35-101t-85-42H675zm0 911v-418h322q50 0 85 41.5t35 100.5v276q0 31 22 53t53 22h98q30 0 52-22t22-53v-276q0-87-24.5-140T1262 896q102-102 102-246q0-145-104-248t-251-103H502q-30 0-52 22t-22 53v1044q0 31 22 53t52 22h98q31 0 53-22t22-53z"/>`),
		g.Group(children),
	)
}