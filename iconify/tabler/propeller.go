package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Propeller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 13a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/><path d="M14.167 10.5c.722-1.538 1.156-3.043 1.303-4.514C15.69 4.356 14.708 3 12 3S8.31 4.357 8.53 5.986c.147 1.471.581 2.976 1.303 4.514m3.336 6.251c.97 1.395 2.057 2.523 3.257 3.386c1.3 1 2.967.833 4.321-1.512c1.354-2.345.67-3.874-.85-4.498c-1.348-.608-2.868-.985-4.562-1.128M8.664 13c-1.693.143-3.213.52-4.56 1.128c-1.522.623-2.206 2.153-.852 4.498s3.02 2.517 4.321 1.512c1.2-.863 2.287-1.991 3.258-3.386"/></g>`),
		g.Group(children),
	)
}