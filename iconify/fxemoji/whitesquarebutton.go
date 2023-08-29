package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whitesquarebutton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#D9DDE8" d="M467 69.76C467 56.142 455.858 45 442.24 45H66.951c-12.256 0-22.284 10.028-22.284 22.284v377.764c0 12.256 10.028 22.284 22.284 22.284H442.24c13.618 0 24.76-11.142 24.76-24.76V69.76z"/><path fill="#132028" d="M371 154.504c0-7.427-6.077-13.504-13.504-13.504H152.82c-6.684 0-12.154 5.469-12.154 12.154V359.18c0 6.684 5.469 12.154 12.154 12.154h204.676c7.427 0 13.504-6.077 13.504-13.504V154.504z"/>`),
		g.Group(children),
	)
}