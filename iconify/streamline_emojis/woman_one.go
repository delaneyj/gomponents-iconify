package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomanOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M11.5 44.5a12.5 1.5 0 1 0 25 0a12.5 1.5 0 1 0-25 0Z" opacity=".15"/><path fill="#724f3d" d="M24 1.5a16.16 16.16 0 0 1 16.16 16.16v21.08A1.81 1.81 0 0 1 39 40.43a40.31 40.31 0 0 1-30 0a1.81 1.81 0 0 1-1.14-1.69V17.66A16.16 16.16 0 0 1 24 1.5Z"/><path fill="#a86c4d" d="M24 1.5A16.16 16.16 0 0 0 7.84 17.66v4.47a16.16 16.16 0 0 1 32.32 0v-4.47A16.16 16.16 0 0 0 24 1.5Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 1.5a16.16 16.16 0 0 1 16.16 16.16v21.08A1.81 1.81 0 0 1 39 40.43h0a40.31 40.31 0 0 1-30 0h0a1.81 1.81 0 0 1-1.14-1.69V17.66A16.16 16.16 0 0 1 24 1.5Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37.47 21A13 13 0 0 1 29 13.57l-.93-2.21a23.56 23.56 0 0 1-14.78 9.17l-2.8.51a2.7 2.7 0 0 0 0 5.39h.17a13.45 13.45 0 0 0 26.6 0h.17a2.7 2.7 0 0 0 0-5.39Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.93 23.18A1.15 1.15 0 1 0 17.07 22a1.14 1.14 0 0 0-1.14 1.18Zm16.14 0A1.15 1.15 0 1 1 30.93 22a1.14 1.14 0 0 1 1.14 1.18Z"/><path fill="#ff6242" d="M20.31 30.47a.57.57 0 0 0-.44.2a.54.54 0 0 0-.13.46a4.32 4.32 0 0 0 8.52 0a.54.54 0 0 0-.13-.46a.57.57 0 0 0-.44-.2Z"/><path fill="#ffa694" d="M24 32.24a5.37 5.37 0 0 0-3.29 1a4.36 4.36 0 0 0 6.58 0a5.37 5.37 0 0 0-3.29-1Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M20.31 30.47a.57.57 0 0 0-.44.2a.54.54 0 0 0-.13.46a4.32 4.32 0 0 0 8.52 0a.54.54 0 0 0-.13-.46a.57.57 0 0 0-.44-.2Z"/><path fill="#ffaa54" d="M13.69 28.06a1.49.89 0 1 0 2.98 0a1.49.89 0 1 0-2.98 0Zm17.64 0a1.49.89 0 1 0 2.98 0a1.49.89 0 1 0-2.98 0Z"/>`),
		g.Group(children),
	)
}