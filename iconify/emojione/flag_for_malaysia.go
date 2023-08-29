package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForMalaysia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ed4c5c" d="M56 14H32v6h27.5c-.9-2.1-2.1-4.2-3.5-6"/><path fill="#f9f9f9" d="M61.4 38c.4-1.9.6-3.9.6-6H2c0 2.1.2 4.1.6 6h58.8"/><path fill="#ed4c5c" d="M32 2v6h18c-5-3.8-11.2-6-18-6"/><path fill="#f9f9f9" d="M32 14h24c-1.7-2.3-3.7-4.3-6-6H32v6m27.5 6H32v6h29.4c-.4-2.1-1.1-4.1-1.9-6"/><path fill="#ed4c5c" d="M32 26v6h30c0-2.1-.2-4.1-.6-6H32M4.5 44h55c.8-1.9 1.5-3.9 1.9-6H2.6c.4 2.1 1.1 4.1 1.9 6"/><path fill="#f9f9f9" d="M8 50h48c1.4-1.8 2.6-3.9 3.5-6h-55c.9 2.1 2.1 4.2 3.5 6"/><path fill="#ed4c5c" d="M8 50c1.7 2.3 3.7 4.3 6 6h36c2.3-1.7 4.3-3.7 6-6H8z"/><path fill="#f9f9f9" d="M14 56c5 3.8 11.2 6 18 6s13-2.2 18-6H14"/><path fill="#2a5f9e" d="M32 2C15.4 2 2 15.4 2 32h30V2z"/><path fill="#ffe62e" d="M19.9 25.6c-3 0-5.5-2.5-5.5-5.6c0-3.1 2.5-5.6 5.5-5.6c1.2 0 2.2.4 3.1 1c-1.3-1.5-3.1-2.4-5.2-2.4c-3.8 0-6.8 3.1-6.8 7s3.1 7 6.8 7c2.1 0 3.9-.9 5.2-2.4c-.9.6-2 1-3.1 1m6.1-6.8l.7-1.8l-.2 1.9l1.4-1.3l-1 1.7l1.8-.6l-1.6 1l1.9.3l-1.9.3l1.6 1l-1.8-.6l1 1.7l-1.4-1.3l.2 1.9l-.7-1.8l-.7 1.8l.2-1.9l-1.4 1.3l1-1.7l-1.8.6l1.6-1L23 20l1.9-.3l-1.6-1l1.8.6l-1-1.7l1.4 1.3l-.2-1.9z"/>`),
		g.Group(children),
	)
}