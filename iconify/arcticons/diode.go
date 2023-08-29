package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.75 43.5h16.5m-8.25 0L13.79 25.8h20.43L24 43.5zm8.25-39h-16.5m8.25 0l10.22 17.69H13.79L24 4.5z"/>`),
		g.Group(children),
	)
}