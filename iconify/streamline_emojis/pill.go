package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M8.5 41.77a15.5 2.02 0 1 0 31 0a15.5 2.02 0 1 0-31 0Z" opacity=".15"/><path fill="#ff6242" d="M8.09 38.88a9.75 9.75 0 0 1 0-13.79l18-18a9.75 9.75 0 0 1 13.82 13.76l-18 18a9.75 9.75 0 0 1-13.82.03Z"/><path fill="#ff866e" d="m12.33 29.34l18-18A9.78 9.78 0 0 1 41.7 9.52a9.86 9.86 0 0 0-1.79-2.46a9.75 9.75 0 0 0-13.79 0l-18 18a9.75 9.75 0 0 0 0 13.79a9.86 9.86 0 0 0 2.46 1.79a9.75 9.75 0 0 1 1.75-11.3Z"/><path fill="#00b8f0" d="m26.12 7.06l-9 9l13.77 13.81l9-9A9.75 9.75 0 1 0 26.12 7.06Z"/><path fill="#4acfff" d="M41.7 9.52a9.86 9.86 0 0 0-1.79-2.46a9.75 9.75 0 0 0-13.79 0l-9 9l4.24 4.24l9-9A9.78 9.78 0 0 1 41.7 9.52Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M39.91 7.06a9.75 9.75 0 0 1 0 13.79l-18 18A9.75 9.75 0 0 1 8.09 25.09l18-18a9.75 9.75 0 0 1 13.82-.03Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="m26.12 7.06l-9 9l13.77 13.81l9-9A9.75 9.75 0 1 0 26.12 7.06Z"/>`),
		g.Group(children),
	)
}