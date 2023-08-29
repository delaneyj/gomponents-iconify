package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreetingCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 384v1664H256V401L1408 41v343h256zm-384-169L739 384h541V215zm256 297H384v1408h1152V512zM960 902q37-34 83-52t98-18q55 0 104 21t85 57t57 85t21 105q0 53-20 102t-58 87l-370 369l-370-369q-38-38-58-87t-20-102q0-56 21-105t57-85t84-57t105-21q51 0 97 18t84 52zm279 296q41-41 41-98q0-29-11-54t-30-45t-44-30t-55-11q-57 0-98 41l-82 82l-82-82q-41-41-98-41q-29 0-54 11t-45 30t-30 44t-11 55q0 57 41 98l279 279l279-279z"/>`),
		g.Group(children),
	)
}