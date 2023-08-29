package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlecallmissed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm246-708q10-10 10-25t-10.5-25t-25-10t-24.5 10l-68 68l-68-68q-10-10-24.5-10t-25 10t-10.5 25t10 25l68 68l-68 68q-10 10-10 25t10.5 25t25 10t24.5-10l68-68l68 68q10 10 24.5 10t25-10t10.5-25t-10-25l-68-68zm10 324H640l-64 64q-64-17-151.5-105T320 448l64-64V256q-24-34-52-53t-42-5l-67 67q-55 55-16 173t150 229t229 150t173-16l67-67q14-14-4.5-46T768 640z"/>`),
		g.Group(children),
	)
}