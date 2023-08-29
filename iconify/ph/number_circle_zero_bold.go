package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleZeroBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm0-144c-28.26 0-48 24.67-48 60s19.74 60 48 60s48-24.67 48-60s-19.74-60-48-60Zm0 96c-23.33 0-24-32.32-24-36s.67-36 24-36s24 32.32 24 36s-.67 36-24 36Z"/>`),
		g.Group(children),
	)
}