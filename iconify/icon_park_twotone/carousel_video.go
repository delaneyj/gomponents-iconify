package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarouselVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCarouselVideo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M11 7h26v34H11z"/><path d="M4 11h7v26H4zm33 0h7v26h-7z"/><path fill="#555" d="m22 20l6 4l-6 4v-8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCarouselVideo0)"/>`),
		g.Group(children),
	)
}