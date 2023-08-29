package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playcircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm0-896q-159 0-271.5 112.5T128 512t112.5 271.5T512 896t271.5-112.5T896 512T783.5 240.5T512 128zM384 288l384 224l-384 224V288z"/>`),
		g.Group(children),
	)
}