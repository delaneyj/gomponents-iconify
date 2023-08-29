package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HurtThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m384 34.3l-128 128h106.7L149.3 375.7L234.7 205H128l89.6-143.4C194.9 45 167.5 34.3 128 34.3C64 34.3 0 77 0 205c0 64 64 192 256 298.7C448 397 512 269 512 205c0-128-64-170.7-128-170.7z"/>`),
		g.Group(children),
	)
}