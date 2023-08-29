package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 20H72a36 36 0 0 0-36 36v128a36 36 0 0 0 36 36l-9.6 12.8a12 12 0 1 0 19.2 14.4L102 220h52l20.4 27.2a12 12 0 0 0 19.2-14.4L184 220a36 36 0 0 0 36-36V56a36 36 0 0 0-36-36ZM60 116V84h56v32Zm80-32h56v32h-56ZM72 44h112a12 12 0 0 1 12 12v4H60v-4a12 12 0 0 1 12-12Zm112 152H72a12 12 0 0 1-12-12v-44h136v44a12 12 0 0 1-12 12Zm-80-28a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm80 0a16 16 0 1 1-16-16a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}