package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M305.7 0H71v512h384V149.3H305.7V0zm64 320v42.7H156.3V320h213.4zM327 0v128h128L327 0z"/>`),
		g.Group(children),
	)
}