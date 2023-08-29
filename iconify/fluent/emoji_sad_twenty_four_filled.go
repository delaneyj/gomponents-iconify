package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiSadTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 1.998c5.524 0 10.002 4.478 10.002 10.002c0 5.523-4.478 10-10.002 10c-5.524.001-10-4.477-10-10C1.999 6.476 6.476 1.998 12 1.998Zm0 11.5A5.984 5.984 0 0 0 7.712 15.3a.75.75 0 1 0 1.072 1.05A4.485 4.485 0 0 1 12 14.996c1.225 0 2.37.49 3.211 1.347a.75.75 0 1 0 1.07-1.052a5.984 5.984 0 0 0-4.28-1.795ZM9 8.75a1.25 1.25 0 1 0 0 2.499A1.25 1.25 0 0 0 9 8.75Zm6 0a1.25 1.25 0 1 0 0 2.499a1.25 1.25 0 0 0 0-2.499Z"/>`),
		g.Group(children),
	)
}