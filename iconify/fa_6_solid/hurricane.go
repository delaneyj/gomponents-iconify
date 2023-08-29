package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hurricane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 208C0 104.4 75.7 18.5 174.9 2.6c9.1-1.4 17.1 6 17.1 15.3v63.3c0 8.4 6.5 15.3 14.7 16.5C307 112.5 384 199 384 303.4c0 103.6-75.7 189.5-174.9 205.4c-9.2 1.5-17.1-5.9-17.1-15.2v-63.4c0-8.4-6.5-15.3-14.7-16.5C77 398.9 0 312.4 0 208zm288 48a96 96 0 1 0-192 0a96 96 0 1 0 192 0zm-96-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64z"/>`),
		g.Group(children),
	)
}