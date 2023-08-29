package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm64-133q98-17 176-80L576 635v256zm-128 0V635L272 811q78 63 176 80zM128 512q0 121 70 221l250-250V134q-137 23-228.5 129.5T128 512zm448-378v349l250 250q70-100 70-221q0-142-91.5-248.5T576 134z"/>`),
		g.Group(children),
	)
}