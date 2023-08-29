package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoundBanknote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#9450e0" d="M0 46.5h64v9H0z"/><path fill="#c28fef" d="M0 8.5h64v38H0z"/><path fill="#94989b" d="M24 46.5h16v9H24z"/><path fill="#9450e0" d="M4 12.5h56v30H4z"/><path fill="#c28fef" d="M7 15.5h50v24H7z"/><circle cx="45" cy="27.5" r="8" fill="#9450e0"/><path fill="#d0d0d0" d="M24 8.5h16v38H24z"/><path fill="#fff" d="M16 30.4c0-.3 0-.6-.1-1h2.9V27h-4.1c-.1-.1-.1-.1-.1-.2c-.1-.2-.3-.4-.4-.5c-.6-.8-.9-1.2-.9-1.9c0-1.5 1.2-2.7 2.8-2.7c1.5 0 2.8 1.2 2.8 2.7H21c0-2.7-2.2-4.9-5-4.9s-5 2.2-5 4.9c0 1.1.5 1.9 1 2.6h-1v2.3h2.5c.2.5.2.8.2 1c0 1.8-1.8 2.9-2.8 3v2.2h10v-2.2h-6c.7-.8 1.1-1.8 1.1-2.9"/>`),
		g.Group(children),
	)
}