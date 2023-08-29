package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeClosedSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.497 6.666C3.56 7.848 5.186 9 7.5 9s3.939-1.152 5.003-2.334a9.37 9.37 0 0 0 1.449-2.164a4.967 4.967 0 0 0 .08-.18l.004-.007v-.001l.464.186l.464.186v.002l-.003.004l-.005.014a3.334 3.334 0 0 1-.1.222a10.37 10.37 0 0 1-1.61 2.406a9.204 9.204 0 0 1-.598.607l1.706 1.705l-.708.708l-1.774-1.775A7.304 7.304 0 0 1 8 9.984V12H7V9.984A7.304 7.304 0 0 1 3.128 8.58l-1.774 1.775l-.708-.708l1.706-1.705a9.237 9.237 0 0 1-.599-.607a10.367 10.367 0 0 1-1.61-2.406a6.064 6.064 0 0 1-.099-.222L.04 4.692l-.002-.004v-.001H.035L.5 4.5l.464-.186l.004.008a2.633 2.633 0 0 0 .08.18a9.368 9.368 0 0 0 1.449 2.164ZM.964 4.314Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}