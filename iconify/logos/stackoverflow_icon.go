package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackoverflowIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#BCBBBB" d="M216.33 276.188v-81.211h26.953v108.165H0V194.977h26.954v81.211z"/><path fill="#F48023" d="m56.708 187.276l132.318 27.654l5.6-26.604L62.31 160.672l-5.601 26.604Zm17.502-63.009l122.517 57.058l11.202-24.503L85.412 99.414L74.21 124.267Zm33.955-60.208l103.964 86.462l17.152-20.653l-103.964-86.462l-17.152 20.653ZM175.375 0L153.67 16.102l80.511 108.515l21.703-16.102L175.375 0ZM53.906 248.884h135.119V221.93H53.907v26.954Z"/>`),
		g.Group(children),
	)
}