package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watchmaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="23.831" cy="24" r="13.771" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.831 14.003V24h-9.997m23.458 2.914h3.97v-5.828H37.29m-5.67-8.441L30.441 4.5h-13.22l-1.179 8.145m0 22.71l1.179 8.145h13.22l1.179-8.145"/>`),
		g.Group(children),
	)
}