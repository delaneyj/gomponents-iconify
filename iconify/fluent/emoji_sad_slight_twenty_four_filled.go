package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiSadSlightTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22.002 12.002C22.002 6.478 17.524 2 12 2C6.476 2 2 6.478 2 12.002c0 5.523 4.477 10.001 10.001 10.001c5.524 0 10.002-4.478 10.002-10.001Zm-14.25-2a1.25 1.25 0 1 1 2.498 0a1.25 1.25 0 0 1-2.499 0Zm6 0a1.25 1.25 0 1 1 2.498 0a1.25 1.25 0 0 1-2.499 0ZM15.75 14h.6a.75.75 0 0 1 0 1.5h-.6c-.618 0-1.337.16-1.998.418c-.669.26-1.197.588-1.472.862a.75.75 0 1 1-1.06-1.06c.475-.476 1.212-.898 1.989-1.2c.784-.306 1.69-.52 2.541-.52Z"/>`),
		g.Group(children),
	)
}