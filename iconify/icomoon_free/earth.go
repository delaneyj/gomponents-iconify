package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Earth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm0 15a6.96 6.96 0 0 1-2.769-.57l3.643-4.098A.503.503 0 0 0 9 10V8.5a.5.5 0 0 0-.5-.5C6.735 8 4.872 6.165 4.854 6.146A.5.5 0 0 0 4.5 6h-2a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .276.447L4 10.809v2.936A6.992 6.992 0 0 1 1 8a6.97 6.97 0 0 1 .674-3H3.5c.133 0 .26-.053.354-.146l2-2A.5.5 0 0 0 6 2.5V1.29A6.989 6.989 0 0 1 8 1c1.1 0 2.141.254 3.067.706A2.98 2.98 0 0 0 10 3.999c0 .801.312 1.555.879 2.121a2.994 2.994 0 0 0 2.268.875c.216.809.605 2.917-.131 5.818a.466.466 0 0 0-.013.082a6.979 6.979 0 0 1-5.002 2.104z"/>`),
		g.Group(children),
	)
}