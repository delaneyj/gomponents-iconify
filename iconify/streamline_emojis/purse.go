package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Purse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#daedf7" d="M19.1 7.6a2.45 2.45 0 1 0 4.9 0a2.45 2.45 0 1 0-4.9 0Z"/><path fill="#daedf7" d="M24 7.6a2.45 2.45 0 1 0 4.9 0a2.45 2.45 0 1 0-4.9 0Z"/><path fill="#fff" d="M21.55 7.56a2.6 2.6 0 0 1 2.23 1.06a2.46 2.46 0 1 0-4.47 0a2.61 2.61 0 0 1 2.24-1.06Z"/><path fill="#fff" d="M26.45 7.56a2.61 2.61 0 0 1 2.24 1.06a2.46 2.46 0 1 0-4.47 0a2.6 2.6 0 0 1 2.23-1.06Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M19.1 7.6a2.45 2.45 0 1 0 4.9 0a2.45 2.45 0 1 0-4.9 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 7.6a2.45 2.45 0 1 0 4.9 0a2.45 2.45 0 1 0-4.9 0Z"/><path fill="#daedf7" d="M3.13 26.58a20.87 16.92 0 1 0 41.74 0a20.87 16.92 0 1 0-41.74 0Z"/><path fill="#fff" d="M24 12c9.63 0 17.74 4.29 20.15 10.13C41.74 15 33.63 9.66 24 9.66S6.26 15 3.85 22.14C6.26 16.3 14.37 12 24 12Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M3.13 26.58a20.87 16.92 0 1 0 41.74 0a20.87 16.92 0 1 0-41.74 0Z"/><path fill="#45413c" d="M2.51 43.78a21.49 2.22 0 1 0 42.98 0a21.49 2.22 0 1 0-42.98 0Z" opacity=".15"/><path fill="#ff6196" d="M46.42 36.77C46.42 43 35.21 43.5 24 43.5S1.58 43 1.58 36.77c0-12.38 10-23.53 22.42-23.53s22.42 11.15 22.42 23.53Z"/><path fill="#ff87af" d="M24 18c11.27 0 20.6 9.23 22.18 20.23a4.3 4.3 0 0 0 .24-1.46c0-12.38-10-23.53-22.42-23.53S1.58 24.39 1.58 36.77a4.3 4.3 0 0 0 .24 1.46C3.4 27.23 12.73 18 24 18Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M46.42 36.77C46.42 43 35.21 43.5 24 43.5S1.58 43 1.58 36.77c0-12.38 10-23.53 22.42-23.53s22.42 11.15 22.42 23.53ZM24 13.24v5.68"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M24 13.24a18.92 18.92 0 0 0-6.36 6.55M24 13.24a18.92 18.92 0 0 1 6.36 6.55"/>`),
		g.Group(children),
	)
}