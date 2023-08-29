package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cherry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCherry0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><circle cx="14" cy="34" r="8" fill="#fff" stroke-linejoin="round"/><circle cx="35" cy="37" r="7" fill="#fff" stroke-linejoin="round"/><path d="M37 10c-2.651.812-8.372 3.014-11.72 6.26C20.255 21.13 19 24.5 18 27m19-17c-1.117 1.318-3.285 4.596-3.956 8.39C32.035 24.078 33 27.5 34 30M30 4l14 12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCherry0)"/>`),
		g.Group(children),
	)
}