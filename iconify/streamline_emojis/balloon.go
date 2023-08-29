package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Balloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M16.1 45.42a7.9 1.58 0 1 0 15.8 0a7.9 1.58 0 1 0-15.8 0Z" opacity=".15"/><path fill="#ff6196" d="M33.31 12.94C33.31 7 29.14 2.88 24 2.88S14.69 7 14.69 12.94c0 5.37 5.16 11.23 8.39 12.27l-.74 1.7a.65.65 0 0 0 .59.91h2.14a.65.65 0 0 0 .59-.91l-.74-1.7c3.23-1.04 8.39-6.9 8.39-12.27Z"/><path fill="#ff87af" d="M24 7.27c4.64 0 8.48 2.85 9.18 7.08a9.6 9.6 0 0 0 .13-1.41C33.31 7 29.14 2.88 24 2.88S14.69 7 14.69 12.94a9.6 9.6 0 0 0 .13 1.41c.7-4.23 4.54-7.08 9.18-7.08Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.31 12.94C33.31 7 29.14 2.88 24 2.88S14.69 7 14.69 12.94c0 5.37 5.16 11.23 8.39 12.27l-.74 1.7a.65.65 0 0 0 .59.91h2.14a.65.65 0 0 0 .59-.91l-.74-1.7c3.23-1.04 8.39-6.9 8.39-12.27ZM24 27.82c0 6.34 8 7.18 11.78 10.35s3.56 6.33 3.56 6.33"/>`),
		g.Group(children),
	)
}