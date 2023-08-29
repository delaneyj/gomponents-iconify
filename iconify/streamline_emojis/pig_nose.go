package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PigNose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M6.3 44.17a17.7 1.83 0 1 0 35.4 0a17.7 1.83 0 1 0-35.4 0Z" opacity=".15"/><path fill="#ff87af" d="M2.33 20.74a21.67 15.26 0 1 0 43.34 0a21.67 15.26 0 1 0-43.34 0Z"/><path fill="#ffb0ca" d="M24 9.15c15.43 0 20.88 5.85 21.58 13.45q.09-.92.09-1.86c0-8.43-5-15.26-21.67-15.26S2.33 12.31 2.33 20.74q0 .95.09 1.86C3.12 15 8.57 9.15 24 9.15Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2.33 20.74a21.67 15.26 0 1 0 43.34 0a21.67 15.26 0 1 0-43.34 0Z"/><path fill="#ff6196" d="M28.23 20.74a4.81 6.55 0 1 0 9.62 0a4.81 6.55 0 1 0-9.62 0Zm-18.08 0a4.81 6.55 0 1 0 9.62 0a4.81 6.55 0 1 0-9.62 0Z"/><path fill="#e0366f" d="M33 17.54c2.23 0 4.11 2.07 4.65 4.88a9.2 9.2 0 0 0 .16-1.68c0-3.62-2.15-6.55-4.81-6.55s-4.8 2.93-4.8 6.55a8.4 8.4 0 0 0 .16 1.68c.58-2.81 2.45-4.88 4.64-4.88Zm-18 0c2.23 0 4.1 2.07 4.64 4.88a8.4 8.4 0 0 0 .16-1.68c0-3.62-2.15-6.55-4.8-6.55s-4.81 2.93-4.81 6.55a9.2 9.2 0 0 0 .16 1.68c.5-2.81 2.38-4.88 4.65-4.88Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.23 20.74a4.81 6.55 0 1 0 9.62 0a4.81 6.55 0 1 0-9.62 0Zm-18.08 0a4.81 6.55 0 1 0 9.62 0a4.81 6.55 0 1 0-9.62 0Z"/>`),
		g.Group(children),
	)
}