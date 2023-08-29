package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glimmer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 0h256v1280h-256V0zM0 2048v-256h1280v256H0zM486 666l180-180l896 896l-180 180l-896-896z"/>`),
		g.Group(children),
	)
}