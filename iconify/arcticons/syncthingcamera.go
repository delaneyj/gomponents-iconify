package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Syncthingcamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.85 19.582a9 9 0 0 0-15.815 5.785m16.495 5.075a8.958 8.958 0 0 0 1.504-4.983a9.054 9.054 0 0 0-.145-1.622m-16.662 6.101a9.008 9.008 0 0 0 11.73 3.624"/><circle cx="25.331" cy="26.226" r="2.389" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.53 19.484a2.389 2.389 0 1 1 0 4.777q-.104 0-.207-.008a2.388 2.388 0 0 1 .207-4.769Zm-2.045 9.854a2.383 2.383 0 1 1-.666.094a2.389 2.389 0 0 1 .666-.094Zm-13.489-4.172a2.389 2.389 0 0 1 .244 4.764a2.433 2.433 0 0 1-.244.013a2.389 2.389 0 1 1 0-4.777Zm2.378 2.168l4.585-.831m4.186-1.831l2.529-1.298m-2.823 4.695l1.324 1.662"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.832 8.775l-2.768 2.832s-7.409.003-10.774.002A1.79 1.79 0 0 0 4.5 13.4v23.223a1.79 1.79 0 0 0 1.79 1.791h35.418a1.79 1.79 0 0 0 1.791-1.793V13.4a1.79 1.79 0 0 0-1.79-1.793H30.935L28.17 8.775Z"/>`),
		g.Group(children),
	)
}