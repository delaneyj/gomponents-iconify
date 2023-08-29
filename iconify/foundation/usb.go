package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M64.603 35.504a2.777 2.777 0 0 0-2.774-2.72c-.047 0-.091.012-.138.014h-2.039V17.774a1.687 1.687 0 0 0-1.684-1.684H41.765v.03a1.681 1.681 0 0 0-1.386 1.654v15.022h-2.052c-.043-.002-.083-.013-.126-.013a2.776 2.776 0 0 0-2.774 2.72h-.036v45.625h.001a2.782 2.782 0 0 0 2.78 2.781l.014-.001h23.643a2.78 2.78 0 0 0 2.78-2.779V35.504h-.006zm-9.938-2.706h-9.329v-11.72h9.329v11.72z"/><path fill="currentColor" d="M47.506 27.933h4.988v2.072h-4.988z"/>`),
		g.Group(children),
	)
}