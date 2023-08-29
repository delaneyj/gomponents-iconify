package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Babylondotjs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.392 6L19.14 4.123L16.01 5.93l3.253 1.878zM4.737 7.807l10.391-6L12 0L1.608 6Zm4.01 6.07L12 15.758l3.252-1.877L12 12Zm10.515-6.07v8.387L12 20.387l-7.263-4.193V7.806L1.608 6.001v12L12 24l10.392-6V6ZM12 8.245l-3.253 1.878v3.757L12 12l3.252 1.879v-3.757z"/>`),
		g.Group(children),
	)
}