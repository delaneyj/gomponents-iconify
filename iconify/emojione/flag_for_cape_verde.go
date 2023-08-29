package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCapeVerde(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#2a5f9e" d="M32 62c9.8 0 18.5-4.7 24-12H8c5.5 7.3 14.2 12 24 12m0-60C15.4 2 2 15.4 2 32h60C62 15.4 48.6 2 32 2z"/><path fill="#fff" d="M61.4 38c.4-1.9.6-3.9.6-6H2c0 2.1.2 4.1.6 6h58.8"/><path fill="#ed4c5c" d="M4.5 44h55c.8-1.9 1.5-3.9 1.9-6H2.6c.4 2.1 1.1 4.1 1.9 6"/><path fill="#fff" d="M4.5 44c.9 2.1 2.1 4.2 3.5 6h48c1.4-1.8 2.6-3.9 3.5-6h-55"/><path fill="#ffce31" d="m25 26.1l1.2.9l-.4-1.5l1.2-1h-1.5L25 23l-.5 1.5H23l1.2 1l-.4 1.5zm0 30l1.2.9l-.4-1.5l1.2-1h-1.5L25 53l-.5 1.5H23l1.2 1l-.4 1.5zm11-20l1.2.9l-.4-1.5l1.2-1h-1.5L36 33l-.5 1.5H34l1.2 1l-.4 1.5zm-18-7l1.2.9l-.4-1.5l1.2-1h-1.5L18 26l-.5 1.5H16l1.2 1l-.4 1.5zm14 0l1.2.9l-.4-1.5l1.2-1h-1.5L32 26l-.5 1.5H30l1.2 1l-.4 1.5zm-19 7l1.2.9l-.4-1.5l1.2-1h-1.5L13 33l-.5 1.5H11l1.2 1l-.4 1.5zm23 12l1.2.9l-.4-1.5l1.2-1h-1.5L36 45l-.5 1.5H34l1.2 1l-.4 1.5zm-23 0l1.2.9l-.4-1.5l1.2-1h-1.5L13 45l-.5 1.5H11l1.2 1l-.4 1.5zm18 6l1.2.9l-.4-1.5l1.2-1h-1.5L31 51l-.5 1.5H29l1.2 1l-.4 1.5zm-13 0l1.2.9l-.4-1.5l1.2-1h-1.5L18 51l-.5 1.5H16l1.2 1l-.4 1.5z"/>`),
		g.Group(children),
	)
}