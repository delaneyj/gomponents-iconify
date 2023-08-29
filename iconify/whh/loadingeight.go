package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loadingeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 0q139 0 257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024t-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0z"/>`),
		g.Group(children),
	)
}