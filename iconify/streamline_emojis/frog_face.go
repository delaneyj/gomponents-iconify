package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrogFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M12.5 44.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/><path fill="#6dd627" d="M37.3 21.87V15.8A7.22 7.22 0 0 0 24 12.17a7.22 7.22 0 0 0-13.28 3.63v6.07c-3.21 1.87-5.2 4.41-5.2 7.22C5.5 34.84 13.78 39.5 24 39.5s18.5-4.66 18.5-10.41c0-2.81-1.99-5.35-5.2-7.22Z"/><path fill="#9ceb60" d="M17.93 11.75a7.16 7.16 0 0 1 4.94 2a1.64 1.64 0 0 0 2.26 0a7.21 7.21 0 0 1 12.15 5V15.8A7.22 7.22 0 0 0 24 12.17a7.22 7.22 0 0 0-13.28 3.63v2.89a7.22 7.22 0 0 1 7.21-6.94Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M37.3 21.87V15.8h0A7.22 7.22 0 0 0 24 12.17a7.22 7.22 0 0 0-13.28 3.63h0v6.07c-3.21 1.87-5.2 4.41-5.2 7.22C5.5 34.84 13.78 39.5 24 39.5s18.5-4.66 18.5-10.41c0-2.81-1.99-5.35-5.2-7.22Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M11.28 30.83S13 34.88 24 34.88s12.72-4.05 12.72-4.05"/><path fill="#c8ffa1" d="M38.5 26.61a1.82 1.82 0 1 1-1.82-1.82a1.81 1.81 0 0 1 1.82 1.82Zm-25.36 0a1.82 1.82 0 1 1-1.82-1.82a1.82 1.82 0 0 1 1.82 1.82Z"/><path fill="#fff" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M20.53 26.78a3.37 3.37 0 0 1 2.31 2.31m4.63-2.31a3.37 3.37 0 0 0-2.31 2.31"/><path fill="#fff" d="M18 11.75a4.34 4.34 0 0 0-4.34 4.34v2.31a4.34 4.34 0 0 0 8.68 0v-2.31A4.34 4.34 0 0 0 18 11.75Z"/><path fill="#f0f0f0" d="M18 11.75a4.34 4.34 0 0 0-4.34 4.34v1.68a4.34 4.34 0 0 1 8.68 0v-1.68A4.34 4.34 0 0 0 18 11.75Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M16.5 17.93a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M18 11.75a4.34 4.34 0 0 0-4.34 4.34v2.31a4.34 4.34 0 0 0 8.68 0v-2.31A4.34 4.34 0 0 0 18 11.75Z"/><path fill="#fff" d="M30 11.75a4.34 4.34 0 0 1 4.34 4.34v2.31a4.34 4.34 0 0 1-8.68 0v-2.31A4.34 4.34 0 0 1 30 11.75Z"/><path fill="#f0f0f0" d="M30 11.75a4.34 4.34 0 0 1 4.34 4.34v1.68a4.34 4.34 0 0 0-8.68 0v-1.68A4.34 4.34 0 0 1 30 11.75Z"/><path fill="#45413c" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M28.5 17.93a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M30 11.75a4.34 4.34 0 0 1 4.34 4.34v2.31a4.34 4.34 0 0 1-8.68 0v-2.31A4.34 4.34 0 0 1 30 11.75Z"/>`),
		g.Group(children),
	)
}