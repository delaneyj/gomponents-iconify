package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm339 507v285h334q49 0 84.5-41.5T1129 650t-35-101t-85-42H675zM428 374v1044q0 31 21.5 53t52.5 22h507q147 0 251-103t104-248q0-144-102-246q102-102 102-246q0-145-104-248t-251-103H502q-31 0-52.5 22T428 374zm581 911H675v-285h334q49 0 84.5 41.5T1129 1142t-35 101t-85 42z"/>`),
		g.Group(children),
	)
}