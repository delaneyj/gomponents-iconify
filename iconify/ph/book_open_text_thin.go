package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpenTextThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 52h-64a36 36 0 0 0-32 19.54A36 36 0 0 0 96 52H32a12 12 0 0 0-12 12v128a12 12 0 0 0 12 12h64a28 28 0 0 1 28 28a4 4 0 0 0 8 0a28 28 0 0 1 28-28h64a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12ZM96 196H32a4 4 0 0 1-4-4V64a4 4 0 0 1 4-4h64a28 28 0 0 1 28 28v121.4A35.94 35.94 0 0 0 96 196Zm132-4a4 4 0 0 1-4 4h-64a35.94 35.94 0 0 0-28 13.41V88a28 28 0 0 1 28-28h64a4 4 0 0 1 4 4Zm-24-96a4 4 0 0 1-4 4h-40a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4Zm0 32a4 4 0 0 1-4 4h-40a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4Zm0 32a4 4 0 0 1-4 4h-40a4 4 0 0 1 0-8h40a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}