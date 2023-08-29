package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Disc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 176a80 80 0 1 0 80 80a80.09 80.09 0 0 0-80-80Zm0 112a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/><path fill="currentColor" d="M414.39 97.61A224 224 0 1 0 97.61 414.39A224 224 0 1 0 414.39 97.61ZM256 368a112 112 0 1 1 112-112a112.12 112.12 0 0 1-112 112Z"/>`),
		g.Group(children),
	)
}