package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckPickup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m368.6 96l76.8 96H288V96h80.6zM224 80v112H64c-17.7 0-32 14.3-32 32v64c-17.7 0-32 14.3-32 32s14.3 32 32 32h33.1c-.7 5.2-1.1 10.6-1.1 16c0 61.9 50.1 112 112 112s112-50.1 112-112c0-5.4-.4-10.8-1.1-16h66.3c-.7 5.2-1.1 10.6-1.1 16c0 61.9 50.1 112 112 112s112-50.1 112-112c0-5.4-.4-10.8-1.1-16h33c17.7 0 32-14.3 32-32s-14.3-32-32-32v-64c0-17.7-14.3-32-32-32h-48.6L418.6 56c-12.1-15.2-30.5-24-50-24H272c-26.5 0-48 21.5-48 48zm0 288a48 48 0 1 1-96 0a48 48 0 1 1 96 0zm288 0a48 48 0 1 1-96 0a48 48 0 1 1 96 0z"/>`),
		g.Group(children),
	)
}