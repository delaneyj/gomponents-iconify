package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHorizontalOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.95 6.952c2.901.15 5.803-.323 8.705-1.42A1 1 0 0 1 21 6.466V17m-3.212.806c-4.483-1.281-8.966-1.074-13.449.622A.993.993 0 0 1 3 17.493V6.466a1 1 0 0 1 1.338-.935c.588.221 1.176.418 1.764.59M3 3l18 18"/>`),
		g.Group(children),
	)
}