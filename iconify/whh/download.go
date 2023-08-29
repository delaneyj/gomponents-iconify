package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm364-512H641V192q0-27-19-45.5T577 128H448q-26 0-45 18.5T384 192v320H148q-15 0-18.5 6t8.5 19l342 345q13 14 32.5 14t32.5-14l342-345q12-13 8-19t-19-6z"/>`),
		g.Group(children),
	)
}