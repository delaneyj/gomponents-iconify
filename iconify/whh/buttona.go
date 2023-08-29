package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buttona(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm311-282L571 229q-21-37-59-37q-37 0-55 30L201 742q-14 22-7 46t30 36t48.5 6t38.5-28l49-98h304l49 98q13 22 38.5 28t48.5-6t30-36t-7-46zM423 576l89-181l89 181H423z"/>`),
		g.Group(children),
	)
}