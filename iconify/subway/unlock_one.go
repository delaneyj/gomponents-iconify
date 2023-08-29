package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnlockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M418.9 232.7h-256v-93.1c0-51.4 41.7-93.1 93.1-93.1s93.1 41.7 93.1 93.1v23.3h46.5v-23.3C395.6 62.5 333.1 0 256 0S116.4 62.5 116.4 139.6v93.1H93.1c-12.8 0-23.3 10.4-23.3 23.3v232.7c0 12.9 10.4 23.3 23.3 23.3h325.8c12.8 0 23.3-10.4 23.3-23.3V256c0-12.9-10.4-23.3-23.3-23.3z"/>`),
		g.Group(children),
	)
}