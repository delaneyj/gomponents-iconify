package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Justgiving(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M31.62 13.235H20.437l-6.525 6.557h8.968c-1.183 2.635-3.823 4.083-6.88 4.083c-4.14 0-7.521-3.563-7.521-7.697c0-4.151 3.381-7.615 7.521-7.615c1.511 0 2.917.391 4.083 1.156l6.053-6.079A15.89 15.89 0 0 0 16 0C7.161 0 0 7.156 0 16c0 8.828 7.161 16 16 16s16-6.817 16-15.651c0-1.193-.136-1.803-.38-3.115z"/>`),
		g.Group(children),
	)
}