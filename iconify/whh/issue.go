package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Issue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-896q-104 0-192.5 51.5t-140 140T128 512t51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512t-51.5-192.5t-140-140T512 128zm32 704h-64q-13 0-22.5-9.5T448 800v-64q0-13 9.5-22.5T480 704h64q13 0 22.5 9.5T576 736v64q0 13-9.5 22.5T544 832zm0-192h-64q-13 0-22.5-9.5T448 608V224q0-13 9.5-22.5T480 192h64q13 0 22.5 9.5T576 224v384q0 13-9.5 22.5T544 640z"/>`),
		g.Group(children),
	)
}