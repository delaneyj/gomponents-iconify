package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourLeaves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 4S33.59 2.201 30 8c-2.672 4.317 1 9 1 9M44 4L31 17M44 4s1.799 10.41-4 14c-4.317 2.672-9-1-9-1m0 0l-3 3m3 11l-3-3m16 16s-10.41 1.799-14-4c-2.672-4.317 1-9 1-9l13 13Zm0 0s1.799-10.41-4-14c-4.317-2.672-9 1-9 1l13 13Zm0 0L31 31l13 13ZM17.264 17l3 3m-16-16s10.41-1.799 14 4c2.672 4.317-1 9-1 9l-13-13Zm0 0s-1.8 10.41 4 14c4.316 2.672 9-1 9-1l-13-13Zm0 0l13 13l-13-13Zm13 27l3-3m-16 16s10.41 1.799 14-4c2.672-4.317-1-9-1-9l-13 13Zm0 0s-1.8-10.41 4-14c4.316-2.672 9 1 9 1l-13 13Zm0 0l13-13l-13 13Z"/>`),
		g.Group(children),
	)
}