package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vignette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 40H40a16 16 0 0 0-16 16v144a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm0 160H40V56h176v144ZM178.05 87.66C164.59 77.56 146.81 72 128 72s-36.59 5.56-50 15.66C63.79 98.27 56 112.6 56 128s7.79 29.73 22 40.34c13.41 10.1 31.19 15.66 50 15.66s36.59-5.56 50.05-15.66C192.21 157.73 200 143.4 200 128s-7.79-29.73-21.95-40.34ZM128 168c-30.88 0-56-17.94-56-40s25.12-40 56-40s56 17.94 56 40s-25.12 40-56 40Z"/>`),
		g.Group(children),
	)
}