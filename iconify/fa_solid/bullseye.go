package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullseye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M248 8C111.03 8 0 119.03 0 256s111.03 248 248 248s248-111.03 248-248S384.97 8 248 8zm0 432c-101.69 0-184-82.29-184-184c0-101.69 82.29-184 184-184c101.69 0 184 82.29 184 184c0 101.69-82.29 184-184 184zm0-312c-70.69 0-128 57.31-128 128s57.31 128 128 128s128-57.31 128-128s-57.31-128-128-128zm0 192c-35.29 0-64-28.71-64-64s28.71-64 64-64s64 28.71 64 64s-28.71 64-64 64z"/>`),
		g.Group(children),
	)
}