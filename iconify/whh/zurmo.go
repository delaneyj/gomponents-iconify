package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zurmo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 896v128L320 896H0V768h640v128H448zM0 576l384-448H0V0h640v128L256 576h384v128H0V576z"/>`),
		g.Group(children),
	)
}