package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyKrwBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 124h-16.17l19.29-47.48a12 12 0 0 0-22.24-9l-23 56.51h-35.8l-23-56.51a12 12 0 0 0-22.24 0L93.92 124H58.08l-23-56.51a12 12 0 0 0-22.24 9L32.17 124H16a12 12 0 0 0 0 24h25.92l23 56.52a12 12 0 0 0 22.24 0l23-56.52H146l23 56.52a12 12 0 0 0 22.24 0l23-56.52H240a12 12 0 0 0 0-24ZM76 168.12L67.83 148h16.34ZM119.83 124l8.17-20.11l8.17 20.11ZM180 168.12L171.83 148h16.34Z"/>`),
		g.Group(children),
	)
}