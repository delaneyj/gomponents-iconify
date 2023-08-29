package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGreece(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#428bc1" d="M56 14H32v6h27.5c-.9-2.1-2.1-4.2-3.5-6"/><path fill="#f9f9f9" d="M32 14h24c-1.7-2.3-3.7-4.3-6-6H32v6m27.5 6H32v6h29.4c-.4-2.1-1.1-4.1-1.9-6"/><path fill="#428bc1" d="M4.5 44h55c.8-1.9 1.5-3.9 1.9-6H2.6c.4 2.1 1.1 4.1 1.9 6"/><path fill="#f9f9f9" d="M8 50h48c1.4-1.8 2.6-3.9 3.5-6h-55c.9 2.1 2.1 4.2 3.5 6"/><path fill="#428bc1" d="M8 50c1.7 2.3 3.7 4.3 6 6h36c2.3-1.7 4.3-3.7 6-6H8z"/><path fill="#f9f9f9" d="M14 56c5 3.8 11.2 6 18 6s13-2.2 18-6H14m6-24V20h12v-6H20V4.5c-2.1.9-4.2 2.1-6 3.5v6H8c-1.4 1.8-2.6 3.9-3.5 6H14v12H2c0 2.1.2 4.1.6 6h58.8c.4-1.9.6-3.9.6-6H20"/><path fill="#428bc1" d="M61.4 26H32v-6H20v12h42c0-2.1-.2-4.1-.6-6M32 2c-4.3 0-8.3.9-12 2.5V14h12V8h18c-5-3.8-11.2-6-18-6M14 14V8c-2.3 1.7-4.3 3.7-6 6h6m-9.5 6C2.9 23.7 2 27.7 2 32h12V20H4.5z"/>`),
		g.Group(children),
	)
}