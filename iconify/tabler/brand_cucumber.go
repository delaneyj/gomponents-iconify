package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCucumber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 10.99c-.01 5.52-4.48 10-10 10.01v-2.26l-.01-.01c-4.28-1.11-6.86-5.47-5.76-9.75a8 8 0 0 1 9.74-5.76C17.5 4.13 20 7.35 20 11v-.01zM10.5 8L10 7m3.5 7l.5 1m-5-2.5L8 13m3 1l-.5 1M13 8l.5-1m2.5 5.5l-1-.5m-6-2l-1-.5"/>`),
		g.Group(children),
	)
}