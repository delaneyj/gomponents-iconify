package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M418.9 232.7h-23.3V69.8c0-38.6-31.3-69.8-69.8-69.8H186.2c-38.6 0-69.8 31.3-69.8 69.8v162.9H93.1c-12.8 0-23.3 10.4-23.3 23.3v232.7c0 12.9 10.4 23.3 23.3 23.3h325.8c12.8 0 23.3-10.4 23.3-23.3V256c0-12.9-10.4-23.3-23.3-23.3zm-69.8 0H162.9V69.8c0-12.9 10.4-23.3 23.3-23.3h139.6c12.8 0 23.3 10.4 23.3 23.3v162.9z"/>`),
		g.Group(children),
	)
}