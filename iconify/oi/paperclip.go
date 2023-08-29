package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M5 0c-.51 0-1.02.21-1.41.59L.81 3.31a2.746 2.746 0 0 0 0 3.88a2.746 2.746 0 0 0 3.88 0l1.25-1.25l-.69-.69l-1.16 1.13l-.09.13c-.69.69-1.81.69-2.5 0c-.68-.68-.66-1.78 0-2.47l2.78-2.75c.39-.39 1.04-.39 1.44 0c.39.39.37 1.01 0 1.41l-2.5 2.47c-.1.1-.27.1-.38 0c-.1-.1-.1-.27 0-.38l.06-.03l.91-.94l-.69-.69l-.97.97c-.48.48-.48 1.27 0 1.75s1.27.49 1.75 0l2.5-2.44c.78-.78.78-2.04 0-2.81C6.01.21 5.51.01 4.99.01z"/>`),
		g.Group(children),
	)
}