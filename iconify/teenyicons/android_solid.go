package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 4a6.473 6.473 0 0 0-2.934.698l-1.65-2.475l-.832.554l1.627 2.44A6.492 6.492 0 0 0 1 10.5V13h13v-2.5a6.492 6.492 0 0 0-2.711-5.282l1.627-2.44l-.832-.555l-1.65 2.475A6.473 6.473 0 0 0 7.5 4ZM5 10H4V9h1v1Zm5 0h1V9h-1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}