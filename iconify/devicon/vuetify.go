package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuetify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#1697f6" d="M65.3 34.414L40.84 76.79L64 116.926l30.672-53.13l30.66-53.128H79Zm0 0"/><path fill="#aeddff" d="m33.34 63.797l1.605 2.793l22.88-39.649l9.402-16.273H2.668Zm0 0"/><path fill="#1867c0" d="M79 10.668C90.594 48.82 64 116.926 64 116.926L40.84 76.789Zm0 0"/><path fill="#7bc6ff" d="M67.227 10.668c-48.844 0-32.282 55.922-32.282 55.922Zm0 0"/>`),
		g.Group(children),
	)
}