package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M478.465 89.022L329.6 47.382L180.3 89.438L41.459 50.052A20 20 0 0 0 16 69.293v340.6a24.093 24.093 0 0 0 17.449 23.089l146.817 41.65l149.365-42.074l140.983 39.436A20 20 0 0 0 496 452.728V112.135a24.08 24.08 0 0 0-17.535-23.113ZM163 436.466L48 403.842V85.17l115 32.624Zm150.615-32.647L195 437.231V118.542L313.615 85.13ZM464 436.91L345.615 403.8V85.089L464 118.2Z"/>`),
		g.Group(children),
	)
}