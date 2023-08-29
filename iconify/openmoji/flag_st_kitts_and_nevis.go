package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagStKittsAndNevis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 17v38l62-38H5z"/><path stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M67 24v-7h-8L5 48v7h8l54-31z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m44.799 33.5l-1.019-5l3.72 3.405l-5-.509l4.37-2.581l-2.071 4.685zm-19 11l-1.019-5l3.72 3.405l-5-.509l4.37-2.581l-2.071 4.685z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}