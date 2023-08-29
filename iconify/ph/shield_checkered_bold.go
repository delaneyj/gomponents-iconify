package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldCheckeredBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 36H48a20 20 0 0 0-20 20v58.77c0 92.37 78.1 123 93.75 128.18a19.82 19.82 0 0 0 5.66 1h1.17a20 20 0 0 0 5.66-1C149.9 237.78 228 207.16 228 114.79V56a20 20 0 0 0-20-20Zm-4 78.79V116h-64V60h64ZM52 60h64v56H52v-1.21Zm2.59 80H116v74.72c-19.67-9.51-52.08-31.35-61.41-74.72ZM140 214.72V140h61.41c-9.33 43.37-41.74 65.21-61.41 74.72Z"/>`),
		g.Group(children),
	)
}