package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleloadersix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512h64q0-91 35.5-174T195 195t143-95.5T512 64V0q104 0 199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zM128 512h64q0 87 43 160.5T351.5 789T512 832t160.5-43T789 672.5T832 512t-43-160.5T672.5 235T512 192v-64q-104 0-192.5 51.5t-140 140T128 512z"/>`),
		g.Group(children),
	)
}