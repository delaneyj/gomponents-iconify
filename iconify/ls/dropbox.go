package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M682 151L552 263l130 113l-217 135l-124-111l-124 111L0 376l130-113L0 151L216 14l125 112L466 14zM140 263l201 127l201-127l-201-127zm201 181l122 108l97-61v65L341 694L122 556v-65l97 61z"/>`),
		g.Group(children),
	)
}