package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M10.5 44.5a13.5 1.5 0 1 0 27 0a13.5 1.5 0 1 0-27 0Z" opacity=".15"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M40.88 22.75a3.27 3.27 0 0 0-2.8-3.23C36.71 12.37 30.93 7 24 7S11.29 12.37 9.92 19.52a3.27 3.27 0 0 0 0 6.47C11.29 33.13 17.07 38.5 24 38.5S36.71 33.13 38.08 26a3.28 3.28 0 0 0 2.8-3.25Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M15.55 21.9a1.2 1.2 0 1 0 1.2-1.2a1.2 1.2 0 0 0-1.2 1.2Zm16.9 0a1.2 1.2 0 1 1-1.2-1.2a1.2 1.2 0 0 1 1.2 1.2Z"/><path fill="#ff6242" d="M19.59 29.2a.69.69 0 0 0-.52.24a.68.68 0 0 0-.16.55a5.16 5.16 0 0 0 10.18 0a.68.68 0 0 0-.16-.55a.69.69 0 0 0-.52-.24Z"/><path fill="#ffa694" d="M24 31.32a6.29 6.29 0 0 0-3.92 1.2a5.21 5.21 0 0 0 7.84 0a6.29 6.29 0 0 0-3.92-1.2Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.59 29.2a.69.69 0 0 0-.52.24a.68.68 0 0 0-.16.55a5.16 5.16 0 0 0 10.18 0a.68.68 0 0 0-.16-.55a.69.69 0 0 0-.52-.24Z"/><path fill="#ffaa54" d="M13.21 27.01a1.55.93 0 1 0 3.1 0a1.55.93 0 1 0-3.1 0Zm18.48 0a1.55.93 0 1 0 3.1 0a1.55.93 0 1 0-3.1 0Z"/><path fill="#724f3d" d="M9.62 19.58a16.41 16.41 0 0 0 5.65-1.75a12.83 12.83 0 0 0 3.63-3v2.72a11.89 11.89 0 0 0 5.51-1.36a8.36 8.36 0 0 0 3.11-2.79a9.87 9.87 0 0 0 4.16 4.21a14.9 14.9 0 0 0 6.4 1.89v-1.57a14.23 14.23 0 0 0-28.46 0Z"/><path fill="#a86c4d" d="M23.85 3.71A14.23 14.23 0 0 0 9.62 17.93v1.65h.13A14.23 14.23 0 0 1 38 19.51h.13v-1.58A14.23 14.23 0 0 0 23.85 3.71Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M9.62 19.58a16.41 16.41 0 0 0 5.65-1.75a12.83 12.83 0 0 0 3.63-3v2.72a11.89 11.89 0 0 0 5.51-1.36a8.36 8.36 0 0 0 3.11-2.79a9.87 9.87 0 0 0 4.16 4.21a14.9 14.9 0 0 0 6.4 1.89v-1.57a14.23 14.23 0 0 0-28.46 0Z"/>`),
		g.Group(children),
	)
}