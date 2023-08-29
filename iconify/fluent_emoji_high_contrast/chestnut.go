package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chestnut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.987 12.586c2.8 2.624 3.52 6.153 2.69 9.258a9.153 9.153 0 0 1-.044.156c-1.09 3.833-4.541 7-9.356 7h-8.569c-4.804 0-8.256-3.167-9.346-7a9.541 9.541 0 0 1-.042-.155c-.82-3.106-.11-6.625 2.69-9.26l9.588-9.032a2.048 2.048 0 0 1 2.8 0l9.589 9.033ZM4.464 22c1.023 2.798 3.647 5 7.244 5h8.57c3.607 0 6.23-2.203 7.254-5H4.464Z"/>`),
		g.Group(children),
	)
}