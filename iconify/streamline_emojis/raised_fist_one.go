package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RaisedFistOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#ffe500" d="M37.29 22.24v-9.16a3.17 3.17 0 0 0-6.34 0v.28a3.74 3.74 0 1 0-7.48 0a3.7 3.7 0 1 0-7.39 0v4.49a3.48 3.48 0 0 0-6.95 0V29a12.71 12.71 0 0 0 12.72 12.7h1.76a15.26 15.26 0 0 0 15.26-15.26a6.31 6.31 0 0 0-1.58-4.2Z"/><path fill="#fff48c" d="M34.12 13.36a3.17 3.17 0 0 1 3.17 3.17v-3.45a3.17 3.17 0 0 0-6.34 0v3.45a3.17 3.17 0 0 1 3.17-3.17ZM12.6 17.79a3.47 3.47 0 0 1 3.48 3.47v-3.45a3.48 3.48 0 0 0-6.95 0v3.45a3.47 3.47 0 0 1 3.47-3.47Z"/><path fill="#fff48c" d="M19.77 13.07a3.7 3.7 0 0 1 3.7 3.7a3.74 3.74 0 0 1 7.48 0v-3.41a3.74 3.74 0 1 0-7.48 0a3.7 3.7 0 1 0-7.39 0v3.45a3.7 3.7 0 0 1 3.69-3.74Z"/><path fill="#45413c" d="M13 45.5a11 1.5 0 1 0 22 0a11 1.5 0 1 0-22 0Z" opacity=".15"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.13 14.34h6.95v11.5H9.13Z"/><path fill="#ffe500" d="M37.3 22.25a6.32 6.32 0 0 0-4.77-2.16H17.32a1.23 1.23 0 0 0-1.24 1.24v.78a5.51 5.51 0 0 0 5.51 5.51h17.24c0-.39.05-.79.05-1.18a6.34 6.34 0 0 0-1.58-4.19Z"/><path fill="#fff48c" d="M17.32 22.24h15.21a6.35 6.35 0 0 1 6 4.24h.36a6.34 6.34 0 0 0-6.34-6.35H17.32a1.23 1.23 0 0 0-1.24 1.24v.78a6 6 0 0 0 .08.95a1.23 1.23 0 0 1 1.16-.86Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M38.87 26.44a6.34 6.34 0 0 0-6.34-6.35H17.32a1.23 1.23 0 0 0-1.24 1.24v.78a5.51 5.51 0 0 0 5.51 5.51h7.54"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M23.47 13.36v6.73"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37.29 22.24v-9.16a3.17 3.17 0 0 0-6.34 0v.28a3.74 3.74 0 1 0-7.48 0v0a3.7 3.7 0 1 0-7.39 0v4.49a3.48 3.48 0 0 0-6.95 0V29a12.71 12.71 0 0 0 12.72 12.7h1.76a15.26 15.26 0 0 0 15.26-15.26a6.31 6.31 0 0 0-1.58-4.2Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30.95 13.36v6.73"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M29 27.62h0a6.12 6.12 0 0 0-6.12 6.12v1.81"/>`),
		g.Group(children),
	)
}