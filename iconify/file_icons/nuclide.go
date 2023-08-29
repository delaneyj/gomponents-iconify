package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nuclide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m375.545 285.492l30.687-30.417V0L254.255 150.057l-8.358-6.791l29.666-29.393V0L177.68 96.646C25.428 11.84-51.87 67.441 38.741 233.83l-22.38 22.099V512l119.85-118.855V251.943l91.71-90.867l8.35 6.739l-89.242 88.114V512l81.773-81.055c179.453 112.072 257.679 34.248 146.743-145.453zM159.093 114.999L57.424 215.383C-11.748 76.538 35.257 53.035 159.094 115zm88.107 297.71l110.085-109.117c89.144 148.933 44.25 198.249-110.085 109.117z"/>`),
		g.Group(children),
	)
}