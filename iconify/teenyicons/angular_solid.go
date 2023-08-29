package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngularSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 4.247L9.142 8H5.858L7.5 4.247Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.363.02a.5.5 0 0 1 .274 0l7 2a.5.5 0 0 1 .36.535l-1 9a.5.5 0 0 1-.273.392l-6 3a.5.5 0 0 1-.448 0l-6-3a.5.5 0 0 1-.273-.392l-1-9a.5.5 0 0 1 .36-.536l7-2ZM7.5 1.752l3.958 9.047l-.916.4L9.579 9H5.421l-.963 2.2l-.916-.4L7.5 1.753Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}