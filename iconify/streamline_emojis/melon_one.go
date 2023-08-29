package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MelonOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M12.6 42.19a13.24 1.81 0 1 0 26.48 0a13.24 1.81 0 1 0-26.48 0Z" opacity=".15"/><path fill="#9ceb60" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m43.41 8.94l-6.58 1.19a14.62 14.62 0 0 1-26.16 8.25l-8.8 3.85A22.16 22.16 0 1 0 43.41 8.94Z"/><path fill="#ffaa54" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M43.3 18.12a19.58 19.58 0 0 0-2.08-8.78l-4.39.79a14.62 14.62 0 0 1-26.16 8.25l-6.61 2.89a19.75 19.75 0 0 0 39.24-3.15Z"/><path fill="#ff8a14" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.51 26.4c7.3-5.19 3.32-16.27 3.32-16.27a14.62 14.62 0 0 1-26.16 8.25S20 36 33.51 26.4Z"/>`),
		g.Group(children),
	)
}