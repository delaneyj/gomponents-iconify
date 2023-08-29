package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyEthLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m220.72 124.29l-88-112a6 6 0 0 0-9.44 0l-88 112a6 6 0 0 0 0 7.42l88 112a6 6 0 0 0 9.44 0l88-112a6 6 0 0 0 0-7.42ZM134 33.35l72.56 92.35l-72.56 33Zm-12 125.33l-72.56-33L122 33.35Zm0 13.18v50.79l-62.08-79Zm12 0l62.08-28.21l-62.08 79Z"/>`),
		g.Group(children),
	)
}