package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LitecoinSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 64c0-17.7-14.3-32-32-32S64 46.3 64 64v149.6l-40.8 11.6c-17 4.9-26.8 22.6-22 39.6s22.6 26.8 39.6 22l23.2-6.7V448c0 17.7 14.3 32 32 32h256c17.7 0 32-14.3 32-32s-14.3-32-32-32H128V261.9l136.8-39.1c17-4.9 26.8-22.6 22-39.6s-22.6-26.8-39.6-22L128 195.3V64z"/>`),
		g.Group(children),
	)
}