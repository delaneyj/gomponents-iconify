package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chestnut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M6.09 44.66a17.91 2.34 0 1 0 35.82 0a17.91 2.34 0 1 0-35.82 0Z" opacity=".15"/><path fill="#bf8256" d="M44.48 22.39C46.77 32 39.29 41.24 28 43.93S5.87 41.76 3.58 32.14C.33 18.53 15 12.87 18.06 5a.76.76 0 0 1 1.21-.29c6.31 5.64 21.99 4.19 25.21 17.68Z"/><path fill="#dea47a" d="M5.9 24.59a23.38 23.38 0 0 1 30.78-7.34a15.07 15.07 0 0 1 7.8 10.07c.08.35.12.69.18 1a15 15 0 0 0-.18-6c-3.22-13.49-18.9-12-25.21-17.73a.76.76 0 0 0-1.21.41c-3 7.67-16.83 13.23-14.73 26a16.19 16.19 0 0 1 2.57-6.41Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M44.48 22.39C46.77 32 39.29 41.24 28 43.93S5.87 41.76 3.58 32.14C.33 18.53 15 12.87 18.06 5a.76.76 0 0 1 1.21-.29c6.31 5.64 21.99 4.19 25.21 17.68Z"/><path fill="#f7e5c6" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M40.8 33.55a11.4 11.4 0 0 1-9.42 1.35a11.37 11.37 0 0 0-10.48 2.49a11.3 11.3 0 0 1-9.32 2.74l-3.72-.58C12.55 44 20.16 45.8 28 43.93c7.51-1.79 13.31-6.48 15.69-12.24Z"/>`),
		g.Group(children),
	)
}