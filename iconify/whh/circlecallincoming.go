package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlecallincoming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm246-708q10-10 10-25t-10.5-25t-25-10t-24.5 10L576 398V288q0-13-9.5-22.5T544 256t-22.5 9.5T512 288v192q0 13 9.5 22.5T544 512h192q13 0 22.5-9.5T768 480t-9.5-22.5T736 448H626zm10 324H640l-64 64q-64-17-151.5-105T320 448l64-64V256q-24-34-52-53t-42-5l-67 67q-55 55-16 173t150 229t229 150t173-16l67-67q14-14-4.5-46T768 640z"/>`),
		g.Group(children),
	)
}