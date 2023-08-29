package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M6.41 40.1a17.87 1.79 0 1 0 35.74 0a17.87 1.79 0 1 0-35.74 0Z" opacity=".15"/><path fill="#fff" d="M2.83 18.61a8.94 12.49 0 1 0 17.88 0a8.94 12.49 0 1 0-17.88 0Z"/><path fill="#e0e0e0" d="M11.77 8.85c4.73 0 8.59 4.68 8.91 10.61v-.85c0-6.9-4-12.49-8.94-12.49s-8.91 5.59-8.91 12.49v.85C3.18 13.53 7 8.85 11.77 8.85Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2.83 18.61a8.94 12.49 0 1 0 17.88 0a8.94 12.49 0 1 0-17.88 0Z"/><path fill="#656769" d="M2.83 18.79a5.91 5.91 0 1 0 11.82 0a5.91 5.91 0 1 0-11.82 0Z"/><path fill="#525252" d="M8.74 22.22A5.91 5.91 0 0 1 3 17.55a5.68 5.68 0 0 0-.14 1.24a5.91 5.91 0 0 0 11.82 0a5.68 5.68 0 0 0-.14-1.24a5.9 5.9 0 0 1-5.8 4.67Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M2.83 18.79a5.91 5.91 0 1 0 11.82 0a5.91 5.91 0 1 0-11.82 0Z"/><path fill="#87898c" d="M7.2 14.92a1.54.84 0 1 0 3.08 0a1.54.84 0 1 0-3.08 0Z"/><path fill="#fff" d="M27.29 18.61a8.94 12.49 0 1 0 17.88 0a8.94 12.49 0 1 0-17.88 0Z"/><path fill="#e0e0e0" d="M36.23 8.85c4.73 0 8.59 4.68 8.91 10.61v-.85c0-6.9-4-12.49-8.94-12.49s-8.94 5.59-8.94 12.49v.85c.38-5.93 4.24-10.61 8.97-10.61Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M27.29 18.61a8.94 12.49 0 1 0 17.88 0a8.94 12.49 0 1 0-17.88 0Z"/><path fill="#656769" d="M27.29 18.79a5.91 5.91 0 1 0 11.82 0a5.91 5.91 0 1 0-11.82 0Z"/><path fill="#525252" d="M33.2 22.22a5.9 5.9 0 0 1-5.77-4.67a5.68 5.68 0 0 0-.14 1.24a5.91 5.91 0 1 0 11.82 0a5.68 5.68 0 0 0-.11-1.24a5.9 5.9 0 0 1-5.8 4.67Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M27.29 18.79a5.91 5.91 0 1 0 11.82 0a5.91 5.91 0 1 0-11.82 0Z"/><path fill="#87898c" d="M31.66 14.92a1.54.84 0 1 0 3.08 0a1.54.84 0 1 0-3.08 0Z"/>`),
		g.Group(children),
	)
}