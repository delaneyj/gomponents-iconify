package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M45.5 45.18C45.5 46.29 35.87 47 24 47s-21.5-.71-21.5-1.82s9.63-2 21.5-2s21.5.89 21.5 2Z" opacity=".15"/><path fill="#656769" d="M3.84 3.84h40.31v40.31H3.84Z"/><path fill="#87898c" d="M41.47 3.84H6.53a2.7 2.7 0 0 0-2.69 2.69v3.86a2.69 2.69 0 0 1 2.69-2.68h34.94a2.69 2.69 0 0 1 2.69 2.68V6.53a2.7 2.7 0 0 0-2.69-2.69Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.84 3.84h40.31v40.31H3.84Z"/><path fill="#e8f4fa" d="M8.55 24a15.45 15.45 0 1 0 30.9 0a15.45 15.45 0 1 0-30.9 0Z"/><path fill="#daedf7" d="M24 12.07a15.45 15.45 0 0 1 15.35 13.69a14.6 14.6 0 0 0 .1-1.76a15.45 15.45 0 0 0-30.9 0a14.6 14.6 0 0 0 .1 1.76A15.45 15.45 0 0 1 24 12.07Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M8.55 24a15.45 15.45 0 1 0 30.9 0a15.45 15.45 0 1 0-30.9 0Z"/><path fill="#f0f0f0" d="M34.08 19.3h10.08v9.4H34.08a1.34 1.34 0 0 1-1.34-1.34v-6.72a1.34 1.34 0 0 1 1.34-1.34Z"/><path fill="#fff" d="M34.08 19.3a1.35 1.35 0 0 0-1.35 1.34v3.44a1.35 1.35 0 0 1 1.35-1.34h10.08V19.3Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M34.08 19.3h10.08v9.4h0h-10.08a1.34 1.34 0 0 1-1.34-1.34v-6.72a1.34 1.34 0 0 1 1.34-1.34Z"/><path fill="#fff" d="m24 24l-12.78 8.69A15.5 15.5 0 0 0 17.4 38Zm12.78-8.7a15.6 15.6 0 0 0-6.17-5.3L24 24Z"/><path fill="#e8f4fa" d="m34.9 16.58l1.88-1.28a15.6 15.6 0 0 0-6.17-5.3l-1.4 3a15.33 15.33 0 0 1 5.69 3.58Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M11.22 32.69A15.5 15.5 0 0 0 17.4 38L24 24ZM36.78 15.3a15.6 15.6 0 0 0-6.17-5.3L24 24Z"/><path fill="#c0dceb" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M18.63 24a5.37 5.37 0 1 0 10.74 0a5.37 5.37 0 1 0-10.74 0Z"/><path fill="#87898c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.98 24a2.02 2.02 0 1 0 4.04 0a2.02 2.02 0 1 0-4.04 0Z"/>`),
		g.Group(children),
	)
}