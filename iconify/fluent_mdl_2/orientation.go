package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orientation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 1664H0V768h1408v896zm-128-768H128v640h1152V896zM128 640H0V512h128v128zm0-256H0V256h128v128zm0-256H0V0h128v128zm768 0H768V0h128v128zm-512 0H256V0h128v128zm256 0H512V0h128v128zm256 256H768V256h128v128zm0 256H768V512h128v128zm731-512q102 102 180 197t132 200t81 226t28 273q0 141-36 272t-103 245t-160 207t-208 160t-245 103t-272 37v-128q123 0 237-32t214-90t182-141t140-181t91-214t32-238q0-133-25-242t-74-204t-120-182t-165-177v293h-128V0h512v128h-293z"/>`),
		g.Group(children),
	)
}