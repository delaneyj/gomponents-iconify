package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropletUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.6 11.998a6.66 6.66 0 0 0-.536-1.12l-4.89-7.26c-.42-.626-1.287-.804-1.936-.398a1.376 1.376 0 0 0-.41.397l-4.893 7.26C4.24 13.715 4.9 17.318 7.502 19.423a7.16 7.16 0 0 0 5.002 1.562M19 22v-6m3 3l-3-3l-3 3"/>`),
		g.Group(children),
	)
}