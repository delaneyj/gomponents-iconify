package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm339 552q208 0 247 3q158 11 224 79q76 78 76 258q0 65-9.5 116t-27.5 87t-44 62t-58.5 41.5t-71 25T929 1236t-92.5 4.5t-100.5-.5h-61V552zM428 374v1044q0 31 22 53t52 22h406q280 0 423-150t143-450q0-594-571-594H502q-30 0-52 22t-22 53z"/>`),
		g.Group(children),
	)
}