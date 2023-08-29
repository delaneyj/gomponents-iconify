package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForPakistan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f9f9f9" d="M17 58V6C8 11.2 2 20.9 2 32s6 20.8 15 26z"/><path fill="#699635" d="M32 2c-5.5 0-10.6 1.5-15 4v52c4.4 2.6 9.5 4 15 4c16.6 0 30-13.4 30-30S48.6 2 32 2z"/><g fill="#fff"><path d="M38 38.1c-6.1 0-11-4.8-11-10.8c0-2.9 1.1-5.4 3-7.4c-4.1 2.1-7 6.4-7 11.3c0 7 5.8 12.7 13 12.7s13-5.7 13-12.7c0-.7-.1-1.4-.2-2c-.9 5.1-5.4 8.9-10.8 8.9"/><path d="m40 19.9l2.5-1.9l-1 3.1l2.5 1.8l-3 .1l-1 3l-1-3l-3-.1l2.5-1.8l-1-3.1z"/></g>`),
		g.Group(children),
	)
}