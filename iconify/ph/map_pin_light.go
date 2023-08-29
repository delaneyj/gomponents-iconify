package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 66a38 38 0 1 0 38 38a38 38 0 0 0-38-38Zm0 64a26 26 0 1 1 26-26a26 26 0 0 1-26 26Zm0-112a86.1 86.1 0 0 0-86 86c0 30.91 14.34 63.74 41.47 94.94a252.32 252.32 0 0 0 41.09 38a6 6 0 0 0 6.88 0a252.32 252.32 0 0 0 41.09-38c27.13-31.2 41.47-64 41.47-94.94a86.1 86.1 0 0 0-86-86Zm0 206.51C113 212.93 54 163.62 54 104a74 74 0 0 1 148 0c0 59.62-59 108.93-74 120.51Z"/>`),
		g.Group(children),
	)
}