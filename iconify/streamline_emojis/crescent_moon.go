package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrescentMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M11.17 45.5a14.5 1.5 0 1 0 29 0a14.5 1.5 0 1 0-29 0Z" opacity=".15"/><path fill="#ffe500" d="M28.3 4.65A15.42 15.42 0 1 1 6.64 26.32a.62.62 0 0 0-1.1.52A18.84 18.84 0 1 0 28.83 3.55a.62.62 0 0 0-.53 1.1Z"/><path fill="#ebcb00" d="M41.05 14.43a19.23 19.23 0 0 1 0 5.36a18.84 18.84 0 0 1-35.2 6.35a.58.58 0 0 0-.31.7a18.84 18.84 0 1 0 35.51-12.41Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.3 4.65A15.42 15.42 0 1 1 6.64 26.32a.62.62 0 0 0-1.1.52A18.84 18.84 0 1 0 28.83 3.55a.62.62 0 0 0-.53 1.1Z"/>`),
		g.Group(children),
	)
}