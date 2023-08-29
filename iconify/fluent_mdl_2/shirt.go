package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shirt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2048 384l-128 512h-256v1024H384V896H128L0 384l768-256q0 53 20 99t55 82t81 55t100 20q53 0 99-20t82-55t55-81t20-100l768 256zm-153 84l-524-175q-24 50-60 90t-82 69t-97 44t-108 16q-56 0-108-15t-97-44t-81-69t-61-91L153 468l75 300h284v1024h1024V768h284l75-300z"/>`),
		g.Group(children),
	)
}