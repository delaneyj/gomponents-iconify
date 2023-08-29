package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Saucillatorforandroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.55 22.52c.56-.29 1.18-1.19 3.11 2.74s5.09 3.88 7.59 0s4.25-4.18 6.59 0s5.45 4.13 7.55 0s4.59-4 6.86 0s5.45 4.32 7.55 0s3.66-2.67 3.66-2.67"/>`),
		g.Group(children),
	)
}