package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyholeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 28a100 100 0 1 0 100 100A100.11 100.11 0 0 0 128 28Zm0 192a92 92 0 1 1 92-92a92.1 92.1 0 0 1-92 92Zm36-108a36 36 0 1 0-59.55 27.22L92.57 169a8 8 0 0 0 7.43 11h56a8 8 0 0 0 7.43-11l-11.88-29.82A36.11 36.11 0 0 0 164 112Zm-21 27.42L156 172h-56l13-32.58a4 4 0 0 0-1.37-4.72a28 28 0 1 1 32.78 0a4 4 0 0 0-1.41 4.72Z"/>`),
		g.Group(children),
	)
}