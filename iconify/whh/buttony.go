package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buttony(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm224-824q-23-12-48.5-6T649 221L512 412L375 221q-13-21-38.5-27t-48.5 6t-30 35t7 43l183 257v233q0 26 19 45t45 19t45-19t19-45V535l183-257q14-20 7-43t-30-35z"/>`),
		g.Group(children),
	)
}