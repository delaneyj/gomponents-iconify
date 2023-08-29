package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M418.4 232.7h-23.3v-93.1C395.1 62.5 332.6 0 255.5 0S115.9 62.5 115.9 139.6v93.1H92.6c-12.8 0-23.3 10.4-23.3 23.3v232.7c0 12.9 10.4 23.3 23.3 23.3h325.8c12.8 0 23.3-10.4 23.3-23.3V256c0-12.9-10.5-23.3-23.3-23.3zm-69.8 0H162.4v-93.1c0-51.4 41.7-93.1 93.1-93.1s93.1 41.7 93.1 93.1v93.1z"/>`),
		g.Group(children),
	)
}