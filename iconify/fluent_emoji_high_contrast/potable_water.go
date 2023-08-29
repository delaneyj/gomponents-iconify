package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PotableWater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18 5a1 1 0 1 0 0 2h1v1h-2a6.5 6.5 0 0 0-6 6.481a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5a2.5 2.5 0 0 1 2-2.45V12h8a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-4V7h1a1 1 0 1 0 0-2h-4ZM7.316 17.1l2.458 8.19a1 1 0 0 0 .958.713h4.497a1 1 0 0 0 .958-.71l2.478-8.208a.845.845 0 0 0-1.615-.498l-.968 3.082a.5.5 0 0 1-.477.35h-5.246a.5.5 0 0 1-.479-.356l-.92-3.057a.858.858 0 0 0-1.644.494Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}