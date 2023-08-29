package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Balloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D22F27" d="m33.976 42.691l.03 6.459l.03 6.459l-5.919-3.807l-5.92-3.807l5.89-2.652z"/><circle cx="45" cy="27" r="23" fill="#EA5A47"/><path fill="#D22F27" d="M60.827 10.549a23.041 23.041 0 0 0-4.361-3.417c5.36 8.847 4.224 20.525-3.418 28.166s-19.32 8.778-28.166 3.418a23.042 23.042 0 0 0 3.417 4.36c8.982 8.982 23.545 8.982 32.528 0c8.982-8.982 8.982-23.545 0-32.527z"/><g fill="none" stroke="#000" stroke-miterlimit="10"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.122" d="m34 47.21l.01 1.94l.03 6.46l-5.92-3.81L22.2 48l5.89-2.66l1.95-.88"/><circle cx="45" cy="27" r="23" stroke-width="2"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.725 65.09c.505.04 1.026 0 1.547-.129c2.704-.664 4.41-3.457 3.812-6.238"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M23.14 58.907a4.875 4.875 0 0 1-.258-1.53c-.024-2.785 2.26-5.13 5.102-5.237"/></g>`),
		g.Group(children),
	)
}