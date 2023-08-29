package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PencilSimpleSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M213.93 210.62a8 8 0 1 1-11.84 10.76l-47.69-52.46l-50.4 50.39A15.86 15.86 0 0 1 92.69 224H48a16 16 0 0 1-16-16v-44.69A15.89 15.89 0 0 1 36.68 152l53.6-53.6l-48.2-53a8 8 0 1 1 11.84-10.78ZM227.32 96L183 140.34a8 8 0 0 1-5.66 2.34h-.19a8 8 0 0 1-5.73-2.62L118.07 81.4a8 8 0 0 1 .26-11L160 28.69a16 16 0 0 1 22.63 0l44.69 44.68a16 16 0 0 1 0 22.63ZM216 84.69L171.32 40l-24 24L192 108.69Z"/>`),
		g.Group(children),
	)
}