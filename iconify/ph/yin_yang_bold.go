package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYangBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20ZM44 128a84.09 84.09 0 0 1 84-84a36 36 0 0 1 0 72a60 60 0 0 0-58.81 71.9A83.73 83.73 0 0 1 44 128Zm84 84a36 36 0 0 1 0-72a60 60 0 0 0 58.81-71.9A83.94 83.94 0 0 1 128 212Zm16-36a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm-32-96a16 16 0 1 1 16 16a16 16 0 0 1-16-16Z"/>`),
		g.Group(children),
	)
}