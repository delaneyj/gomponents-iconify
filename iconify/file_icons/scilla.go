package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scilla(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M72.633 93.247L0 225.237l111.618 193.454l81.363-.103l-33.863-58.505l155.886-1.091s37.973-67.69 38.062-67.91c.127-.318-234.121-.398-234.121-.398l-38.32-66.199l73.06-130.731z"/><path fill="currentColor" d="M432.613 289.149c-.726 1.629-74.62 129.126-74.62 129.49c0 .214 81.364.061 81.364.061L512 287.456L400.537 93.247h-81.505l34.429 58.275s-155.103.939-156.241 1.066l-38.317 68.196l235.293.022z"/>`),
		g.Group(children),
	)
}