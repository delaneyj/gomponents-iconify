package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilCan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M320 128c17.7 0 32-14.3 32-32s-14.3-32-32-32H192c-17.7 0-32 14.3-32 32s14.3 32 32 32h32v32H48c-26.5 0-48 21.5-48 48v64.8c0 19 11.2 36.2 28.5 43.9l67.5 30V368c0 26.5 21.5 48 48 48h259.1c18.4 0 35.8-7.9 48-21.7l182.4-206.6c12.3-13.9-.3-35.4-18.4-31.5L448 192l-50.5-25.2c-8.9-4.4-18.7-6.8-28.6-6.8H288v-32h32zM96 208v86.1l-48-21.3V208h48z"/>`),
		g.Group(children),
	)
}