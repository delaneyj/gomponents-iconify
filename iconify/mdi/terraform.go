package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terraform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 5.4v6.4L15.4 15V8.7L21 5.4m-6.2 3.3V15l-5.6-3.2V5.4l5.6 3.3m0 7v6.4l-5.6-3.2v-6.4l5.6 3.2M8.6 5.1v6.4L3 8.3V1.9l5.6 3.2Z"/>`),
		g.Group(children),
	)
}