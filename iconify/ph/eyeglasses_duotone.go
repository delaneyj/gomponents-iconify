package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeglassesDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M104 164a36 36 0 1 1-36-36a36 36 0 0 1 36 36Zm84-36a36 36 0 1 0 36 36a36 36 0 0 0-36-36Z" opacity=".2"/><path d="M200 40a8 8 0 0 0 0 16a16 16 0 0 1 16 16v58.08A44 44 0 0 0 145.68 152h-35.36A44 44 0 0 0 40 130.08V72a16 16 0 0 1 16-16a8 8 0 0 0 0-16a32 32 0 0 0-32 32v92a44 44 0 0 0 87.81 4h32.38a44 44 0 0 0 87.81-4V72a32 32 0 0 0-32-32ZM68 192a28 28 0 1 1 28-28a28 28 0 0 1-28 28Zm120 0a28 28 0 1 1 28-28a28 28 0 0 1-28 28Z"/></g>`),
		g.Group(children),
	)
}