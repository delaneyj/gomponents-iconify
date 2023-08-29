package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M58.2 395.6C26 395.6 0 421.7 0 453.8S26 512 58.2 512c32.1 0 58.2-26 58.2-58.2s-26.1-58.2-58.2-58.2zM0 0v93.1c231.4 0 418.9 187.5 418.9 418.9H512C512 229.2 282.8 0 0 0zm0 186.2v93.1c128.5 0 232.7 104.2 232.7 232.7h93.1c0-180-145.9-325.8-325.8-325.8z"/>`),
		g.Group(children),
	)
}