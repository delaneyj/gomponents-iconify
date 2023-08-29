package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PawOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6.5 3V2a1.5 1.5 0 1 0-3 0v1a1.5 1.5 0 1 0 3 0Zm5 0V2a1.5 1.5 0 0 0-3 0v1a1.5 1.5 0 1 0 3 0Zm3 4.5V7a1.5 1.5 0 0 0-3 0v.5a1.5 1.5 0 0 0 3 0Zm-11 0V7a1.5 1.5 0 1 0-3 0v.5a1.5 1.5 0 1 0 3 0Zm-.645 4.14l2.918-3.543a2.237 2.237 0 0 1 3.454 0l2.918 3.543c.939 1.14.128 2.86-1.35 2.86c-.194 0-.385-.045-.559-.132l-.36-.18a5.315 5.315 0 0 0-4.753 0l-.36.18a1.248 1.248 0 0 1-.558.132c-1.478 0-2.289-1.72-1.35-2.86Z"/>`),
		g.Group(children),
	)
}