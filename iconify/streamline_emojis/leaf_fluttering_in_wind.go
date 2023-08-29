package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeafFlutteringInWind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M16.84 44.49a12 1.5 0 1 0 24 0a12 1.5 0 1 0-24 0Z" opacity=".15"/><path fill="#6dd627" d="M1 7.41a1 1 0 0 1 .57-1c2.87-1.32 15.19-6.26 26.22.82c8.21 5.24 12.67 7.07 14.83 7.71a1 1 0 0 1 .38 1.65C39.68 20 29.65 28.4 16.4 24.47S1.29 10.19 1 7.41Z"/><path fill="#9ceb60" d="M27.79 11.24c6.57 4.21 10.76 6.22 13.28 7.17c.8-.68 1.46-1.31 2-1.82a1 1 0 0 0-.42-1.65c-2.19-.64-6.65-2.47-14.86-7.71C16.76.15 4.44 5.09 1.58 6.41a1 1 0 0 0-.57 1a16.79 16.79 0 0 0 .59 3c2.9-1.34 15.18-6.24 26.19.83Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M1 7.41a1 1 0 0 1 .57-1c2.87-1.32 15.19-6.26 26.22.82c8.21 5.24 12.67 7.07 14.83 7.71a1 1 0 0 1 .38 1.65C39.68 20 29.65 28.4 16.4 24.47S1.29 10.19 1 7.41Z"/><path fill="#6dd627" d="M33.38 32.09A16.19 16.19 0 0 0 26 42.32a1 1 0 0 0 .71 1.16a16.16 16.16 0 0 0 19.84-12.06a1 1 0 0 0-.7-1.16a16.15 16.15 0 0 0-12.47 1.83Z"/><path fill="#9ceb60" d="M26.73 43.48h.13A16.12 16.12 0 0 1 45.72 34a16.73 16.73 0 0 0 .83-2.57a1 1 0 0 0-.7-1.16A16.15 16.15 0 0 0 26 42.32a1 1 0 0 0 .73 1.16Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M33.38 32.09A16.19 16.19 0 0 0 26 42.32a1 1 0 0 0 .71 1.16a16.16 16.16 0 0 0 19.84-12.06a1 1 0 0 0-.7-1.16a16.15 16.15 0 0 0-12.47 1.83Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M46.49 30.79a15.39 15.39 0 0 0-11.11 5.38M1.05 7c5.84-.18 13.58 1.21 19 4.49"/>`),
		g.Group(children),
	)
}