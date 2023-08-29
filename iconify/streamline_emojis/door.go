package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M9.09 45.14a14.91 1.86 0 1 0 29.82 0a14.91 1.86 0 1 0-29.82 0Z" opacity=".15"/><path fill="#debb7e" d="M12.2 3.51h3.73v41H12.2a1.24 1.24 0 0 1-1.2-1.25V4.75a1.24 1.24 0 0 1 1.2-1.24Z"/><path fill="#debb7e" d="M32.07 3.51v4.97H15.92V3.51z"/><path fill="#f0d5a8" d="M15.93 3.51h16.15v2.48H15.93z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M32.07 3.51v4.97H15.92V3.51z"/><path fill="#debb7e" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M32.08 23.39v4.97H15.93v-4.97zm0 16.15v4.97H15.93v-4.97z"/><path fill="#debb7e" d="M32.07 3.51h3.73A1.24 1.24 0 0 1 37 4.75v38.51a1.24 1.24 0 0 1-1.2 1.24h-3.73v-41v.01Z"/><path fill="#f0d5a8" d="M12.2 3.51A1.23 1.23 0 0 0 11 4.75v3.73A2.48 2.48 0 0 1 13.44 6h2.49V3.51Zm23.6 0h-3.73V6h2.49A2.48 2.48 0 0 1 37 8.48V4.75a1.23 1.23 0 0 0-1.2-1.24Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M12.2 3.51h3.73v41h0h-3.73a1.24 1.24 0 0 1-1.2-1.25V4.75a1.24 1.24 0 0 1 1.2-1.24Zm19.87 0h3.73A1.24 1.24 0 0 1 37 4.75v38.51a1.24 1.24 0 0 1-1.2 1.24h-3.73h0v-41v.01Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.56 21.52H37v6.83h0h-1.2a1.24 1.24 0 0 1-1.24-1.24v-5.59h0Z"/><path fill="#ffe500" d="M33.94 22.76a1.86 1.86 0 1 0 3.72 0a1.86 1.86 0 1 0-3.72 0Z"/><path fill="#fff48c" d="M35.8 22.48a1.84 1.84 0 0 1 1.68 1.07a1.76 1.76 0 0 0 .18-.79a1.86 1.86 0 1 0-3.72 0a1.91 1.91 0 0 0 .18.79a1.87 1.87 0 0 1 1.68-1.07Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.94 22.76a1.86 1.86 0 1 0 3.72 0a1.86 1.86 0 1 0-3.72 0Z"/><path fill="#debb7e" d="M15.93 8.48h16.15v14.91H15.93z"/><path fill="#b89558" d="M15.93 8.48h16.15v2.43H15.93z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.93 8.48h16.15v14.91H15.93z"/><path fill="#debb7e" d="M15.93 28.35h16.15v11.18H15.93z"/><path fill="#b89558" d="M15.93 28.35h16.15v2.43H15.93z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.93 28.35h16.15v11.18H15.93z"/>`),
		g.Group(children),
	)
}