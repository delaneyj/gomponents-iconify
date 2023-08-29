package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraRear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M107 427v-43l64 64l-64 64v-43H0v-42h107zm85 0h107v42H192v-42zM256 0q18 0 30.5 12.5T299 43v298q0 18-12.5 30.5T256 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h213zM149.5 128q17.5 0 30-12.5t12.5-30t-12.5-30t-30-12.5t-30 12.5t-12.5 30t12.5 30t30 12.5z"/>`),
		g.Group(children),
	)
}