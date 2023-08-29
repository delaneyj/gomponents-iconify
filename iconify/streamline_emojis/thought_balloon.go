package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThoughtBalloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M14 43.5a13.5 1.5 0 1 0 27 0a13.5 1.5 0 1 0-27 0Z" opacity=".15"/><path fill="#fff" d="M47.5 17.69A13.69 13.69 0 0 0 21.87 11a10.75 10.75 0 1 0-6.36 19.85a10.25 10.25 0 0 0 18.78.51A13.68 13.68 0 0 0 47.5 17.69Z"/><path fill="#f0f0f0" d="M34.65 27.17A10.73 10.73 0 0 1 15 26.65A11.26 11.26 0 0 1 5.46 20v.13a10.74 10.74 0 0 0 10.05 10.72a10.25 10.25 0 0 0 18.78.51A13.68 13.68 0 0 0 47.48 18.1a14.32 14.32 0 0 1-12.83 9.07Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M47.5 17.69A13.69 13.69 0 0 0 21.87 11a10.75 10.75 0 1 0-6.36 19.85a10.25 10.25 0 0 0 18.78.51A13.68 13.68 0 0 0 47.5 17.69Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M.5 41.11a1.96 1.96 0 1 0 3.92 0a1.96 1.96 0 1 0-3.92 0Z"/><path fill="#fff" d="M7.87 37.58a3.42 3.42 0 1 0 6.84 0a3.42 3.42 0 1 0-6.84 0Z"/><path fill="#f0f0f0" d="M11.29 39.53a4.39 4.39 0 0 1-3.39-1.62a3.4 3.4 0 0 0 6.78 0a4.39 4.39 0 0 1-3.39 1.62Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M7.87 37.58a3.42 3.42 0 1 0 6.84 0a3.42 3.42 0 1 0-6.84 0Z"/>`),
		g.Group(children),
	)
}