package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bancircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1025q-104 0-199-40.5t-163.5-109T40.5 712T0 513t40.5-199t109-163.5T313 41T512 0t199 41t163.5 109.5t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zM128 513q0 104 51.5 192.5t140 140T512 897q122 0 222-71L199 291q-71 100-71 222zm384-384q-121 0-221 70l535 535q70-100 70-221q0-105-51.5-193.5T704.5 180T512 129z"/>`),
		g.Group(children),
	)
}