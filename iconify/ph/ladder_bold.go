package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LadderBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M219.27 211.9L168.77 73l7.63-21h7.6a12 12 0 0 0 0-24H88a12 12 0 0 0-1.11 23.94l-58.16 160a12 12 0 0 0 7.17 15.37a11.83 11.83 0 0 0 4.1.69a12 12 0 0 0 11.28-7.9L60 196h38.5l-5.78 15.9a12 12 0 0 0 7.18 15.37a11.79 11.79 0 0 0 4.1.73a12 12 0 0 0 11.28-7.9l17.49-48.1h46.46l17.49 48.1A12 12 0 0 0 208 228a11.83 11.83 0 0 0 4.1-.73a12 12 0 0 0 7.17-15.37ZM124.69 124H86.22L95 100h38.46Zm17.45-48h-38.46l8.73-24h38.46Zm-73.37 96l8.73-24H116l-8.73 24Zm72.73-24l14.5-39.89L170.5 148Z"/>`),
		g.Group(children),
	)
}