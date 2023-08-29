package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fantom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m512 421.827l-160.302-5.58l-106.53 56.309l-40.754-24.82l-130.675 23.13C-7.426 203.02-32.497 78.272 54.614 51.238C133.667 26.704 218.79 40.349 265.46 93.107c40.815 46.139 212.264 266.032 246.54 328.72zM165.912 252.163s27.818-95.954 22.678-130.5l-84.693 86.783l-77.374-9.41l50.227 57.416l29.896-27.27l59.266 22.98z"/>`),
		g.Group(children),
	)
}