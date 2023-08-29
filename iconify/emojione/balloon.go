package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Balloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M52 22.5c0 11.3-8.8 24.4-19.8 24.4c-10.9 0-19.8-13-19.8-24.4C12.5 11.2 21.3 2 32.2 2C43.2 2 52 11.2 52 22.5"/><path fill="#94989b" d="M31.5 49.1V49v.1m.1.1c-.1 0-.1-.1 0 0"/><path fill="#b2c1c0" d="M33 49.2h-1.5c0 1.8-.4 3.9-1.9 5.2c-2.1 1.8-4.9.7-7.4.6c-3-.1-5.6 1.2-7.7 3.2c-.8.7-1.6 1.7-2.3 2.5c-.8.9.3 1.9 1.1 1c1.5-1.8 2.6-3.1 4.6-4.3c2.6-1.5 5-.7 7.8-.6c3.2.2 6.1-1.4 7-4.7c.2-.7.3-2 .3-2.9"/><path fill="#94989b" d="M31.5 49c0-.1 0 0 0 0"/><path fill="#b2c1c0" d="M30.8 48h2.9c.9 0 .9-1.3 0-1.3h-2.9c-.9.1-.9 1.3 0 1.3"/><path fill="#ed4c5c" d="M30.1 50h4.3c1.4 0 1.4-1.9 0-1.9h-4.3c-1.4-.1-1.4 1.9 0 1.9"/>`),
		g.Group(children),
	)
}