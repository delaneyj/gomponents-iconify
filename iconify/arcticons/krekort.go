package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Krekort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="19.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.5 30.729V16.902h2.217a6.913 6.913 0 0 1 0 13.825Zm11.663-13.461v13.825m7.337.002l-5.627-6.913L33.5 17.27m-5.631 6.91h-1.706"/>`),
		g.Group(children),
	)
}