package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Batteryaltcharging(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 1024H128q-53 0-90.5-37.5T0 896V256q0-53 37.5-90.5T128 128h192V32q0-13 9.5-22.5T352 0h192q13 0 22.5 9.5T576 32v96h192q53 0 90.5 37.5T896 256v640q0 53-37.5 90.5T768 1024zm0-709q0-25-17-42t-41-17H186q-24 0-41 17t-17 42v522q0 25 17 42t41 17h524q24 0 41-17t17-42V315zM448 832L320 512h128V320l128 320H448v192z"/>`),
		g.Group(children),
	)
}