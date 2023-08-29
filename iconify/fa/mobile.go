package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M464 1152q0-33-23.5-56.5T384 1072t-56.5 23.5T304 1152t23.5 56.5T384 1232t56.5-23.5T464 1152zm208-160V288q0-13-9.5-22.5T640 256H128q-13 0-22.5 9.5T96 288v704q0 13 9.5 22.5t22.5 9.5h512q13 0 22.5-9.5T672 992zM480 144q0-16-16-16H304q-16 0-16 16t16 16h160q16 0 16-16zm288-16v1024q0 52-38 90t-90 38H128q-52 0-90-38t-38-90V128q0-52 38-90t90-38h512q52 0 90 38t38 90z"/>`),
		g.Group(children),
	)
}