package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M385 445v-63q28-9 46-36.5t18-58.5V159q0-32-18-59.5T385 62V0q124 12 213 107t104 229Q560 436 385 445zm-32.5-126q-13.5 0-22.5-9.5t-9-22.5V159q0-14 9-23t22.5-9t23 9t9.5 23v128q0 13-9.5 22.5t-23 9.5zM321 382v63Q145 436 3 336q15-134 104.5-229T321 0v62q-28 10-46.5 37.5T256 159v128q0 31 18.5 58t46.5 37zm384 22v235q0 105-47 193.5t-128.5 140t-177 51.5t-177-51.5T47 832.5T0 639V404q161 107 352.5 107T705 404z"/>`),
		g.Group(children),
	)
}