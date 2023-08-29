package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markerleft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 768q-125 0-225-73T276 507L0 384l276-123q39-115 139-188T640 0q104 0 192.5 51.5t140 140T1024 384t-51.5 192.5t-140 140T640 768zm0-640q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75z"/>`),
		g.Group(children),
	)
}