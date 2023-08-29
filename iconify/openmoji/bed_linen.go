package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedLinen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M46.355 25.522a78.453 78.453 0 0 0-20.71 0a19.993 19.993 0 0 0 0-13.463a81.964 81.964 0 0 0 20.71 0a34.84 34.84 0 0 0 0 13.463Z"/><path fill="#d0cfce" d="M46.355 25.522a4.3 4.3 0 0 0-4.66 0a28.096 28.096 0 0 0 0-13.463a4.463 4.463 0 0 0 4.66 0q-.295 6.732 0 13.463Z"/><path fill="#fff" d="M50.306 61.536a65.535 65.535 0 0 0-28.611 0a75.3 75.3 0 0 0 0-31.001a68.423 68.423 0 0 0 28.61 0a132.91 132.91 0 0 0 0 31.001Z"/><path fill="#d0cfce" d="M50.306 61.536q-3.093-3.16-6.44 0a106.836 106.836 0 0 0 0-31.001c2.076 2.061 4.226 1.97 6.44 0q-.409 15.5 0 31.001Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M46.355 25.522a78.453 78.453 0 0 0-20.71 0a19.993 19.993 0 0 0 0-13.463a81.964 81.964 0 0 0 20.71 0a34.84 34.84 0 0 0 0 13.463Zm3.951 36.014a65.535 65.535 0 0 0-28.611 0a75.3 75.3 0 0 0 0-31.001a68.423 68.423 0 0 0 28.61 0a132.91 132.91 0 0 0 0 31.001Z"/>`),
		g.Group(children),
	)
}