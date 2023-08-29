package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCinemaFourD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.65 6.956a5.39 5.39 0 0 0 7.494 7.495"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M17.7 12.137A5.738 5.738 0 1 1 11.963 6.4"/><path d="M17.7 12.338v-1.175c0-.47.171-.92.476-1.253a1.56 1.56 0 0 1 1.149-.52c.827 0 1.523.676 1.62 1.573c.037.344.055.69.055 1.037m-9.338-5.6h1.175c.47 0 .92-.176 1.253-.49c.333-.314.52-.74.52-1.184c0-.852-.676-1.57-1.573-1.67A9.496 9.496 0 0 0 12 3"/></g>`),
		g.Group(children),
	)
}