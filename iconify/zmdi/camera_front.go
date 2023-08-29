package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M107 427v-43l64 64l-64 64v-43H0v-42h107zm85 0h107v42H192v-42zm-42.5-256q-17.5 0-30-12.5T107 128t12.5-30.5t30-12.5t30 12.5T192 128t-12.5 30.5t-30 12.5zM256 0q18 0 30.5 12.5T299 43v298q0 18-12.5 30.5T256 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h213zM43 43v224q0-24 36.5-39t70-15t70 15t36.5 39V43H43z"/>`),
		g.Group(children),
	)
}