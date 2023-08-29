package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M468 98q0 30-25.5 54.5T377 188l-43 44v-37h-10q-59 0-101.5-28.5T180 98t42.5-68T324 2t101.5 28T468 98zm-323 93q-57 0-97.5 40T7 327t40.5 95.5T145 462h65v-72h-40q-1-1-7.5 0t-17.5-2.5t-21.5-8.5t-18.5-19t-8-33q0-31 18.5-46.5T152 264h58v-72h-1l-1-1h-63z"/>`),
		g.Group(children),
	)
}