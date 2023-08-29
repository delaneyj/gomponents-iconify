package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NuxtjsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 2a.5.5 0 0 1 .429.243l1.527 2.545l.613-1.042a.5.5 0 0 1 .862 0l5 8.5A.5.5 0 0 1 14.5 13H.5a.5.5 0 0 1-.429-.757l6-10A.5.5 0 0 1 6.5 2ZM5.374 12h6.243L8.465 6.746L5.375 12ZM7.88 5.77L4.214 12h-2.83L6.5 3.472L7.879 5.77Zm1.163-.005L12.783 12h.843L9.5 4.986l-.458.779Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}