package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortcakeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M5.5 44.32a18.5 1.68 0 1 0 37 0a18.5 1.68 0 1 0-37 0Z" opacity=".15"/><path fill="#bf8256" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M40.26 26.68H6.62v15.7a1.12 1.12 0 0 0 1.12 1.12h31.4a1.12 1.12 0 0 0 1.12-1.12Z"/><path fill="#ffe3cf" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M6.62 33.41h33.64v5.61H6.62z"/><path fill="#ff6242" d="M21.2 31.73a1.68 1.68 0 0 1 3.36 0v.56a1.12 1.12 0 1 0 2.24 0v-.56a.56.56 0 1 1 1.12 0v1.56a2.33 2.33 0 0 0 1.87 2.33a2.25 2.25 0 0 0 2.62-2.21v-2.8a1.68 1.68 0 0 0 3.36 0v-.56a1.13 1.13 0 1 1 2.25 0v1.12a3.37 3.37 0 0 0 1.73 2.94a1.12 1.12 0 0 0 1.63-1v-6.43c0-3.3-6.27-8.18-14.86-9.21a3.35 3.35 0 0 0-1.79.29L6.8 26.08a2.25 2.25 0 0 0-1.3 2v8a.57.57 0 0 0 .67.56a3.37 3.37 0 0 0 2.69-3.19a1.79 1.79 0 0 1 1.35-1.76a1.69 1.69 0 0 1 2 1.65V34a1.68 1.68 0 0 0 3.36 0v-1.15a.56.56 0 1 1 1.12 0v1.56a2.33 2.33 0 0 0 1.87 2.33a2.25 2.25 0 0 0 2.62-2.21Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.2 31.73a1.68 1.68 0 0 1 3.36 0v.56a1.12 1.12 0 1 0 2.24 0v-.56a.56.56 0 1 1 1.12 0v1.56a2.33 2.33 0 0 0 1.87 2.33a2.25 2.25 0 0 0 2.62-2.21v-2.8a1.68 1.68 0 0 0 3.36 0v-.56a1.13 1.13 0 1 1 2.25 0v1.12a3.37 3.37 0 0 0 1.73 2.94a1.12 1.12 0 0 0 1.63-1v-6.43c0-3.3-6.27-8.18-14.86-9.21a3.35 3.35 0 0 0-1.79.29L6.8 26.08a2.25 2.25 0 0 0-1.3 2v8a.57.57 0 0 0 .67.56a3.37 3.37 0 0 0 2.69-3.19a1.79 1.79 0 0 1 1.35-1.76a1.69 1.69 0 0 1 2 1.65V34a1.68 1.68 0 0 0 3.36 0v-1.15a.56.56 0 1 1 1.12 0v1.56a2.33 2.33 0 0 0 1.87 2.33a2.25 2.25 0 0 0 2.62-2.21Z"/><path fill="none" stroke="#e04122" stroke-linecap="round" stroke-linejoin="round" d="M11.11 26.68h26.91"/><path fill="#fffef2" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M31.29 22.2v-9a1.12 1.12 0 0 0-1.12-1.12h-2.25a1.12 1.12 0 0 0-1.12 1.12v9a2.43 2.43 0 0 0 2.25 1.12a2.43 2.43 0 0 0 2.24-1.12Z"/><path fill="#ffe500" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M26.47 6.56a2.66 2.66 0 1 0 5.18 1.23c.46-1.94-2.5-2.26-2.43-5.14a5.63 5.63 0 0 0-2.75 3.91Z"/>`),
		g.Group(children),
	)
}