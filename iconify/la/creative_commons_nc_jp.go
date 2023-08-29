package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsNcJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.065 0 11 4.935 11 11c0 1.238-.217 2.425-.596 3.537l-8.851-3.941L21 11h-2.375L16 14.313L13.531 11h-2.375l1.969 2.625l-6.717-2.99C8.294 7.278 11.884 5 16 5zM5.596 12.463L13 15.76V18h2v1h-2v2h2v1h2v-1h2v-2h-2v-1h1.033l7.559 3.365C23.706 24.722 20.116 27 16 27C9.935 27 5 22.065 5 16c0-1.238.217-2.425.596-3.537z"/>`),
		g.Group(children),
	)
}