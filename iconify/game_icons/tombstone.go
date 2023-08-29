package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tombstone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 46c-81.822 0-150 63.366-150 150v150c0 6.668-.757 23.558 0 30h300c.757-6.442 0-23.332 0-30V196c0-86.634-68.178-150-150-150zM16 406v60h480v-60H16z"/>`),
		g.Group(children),
	)
}