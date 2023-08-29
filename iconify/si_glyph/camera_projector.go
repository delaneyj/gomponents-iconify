package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraProjector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.997 5.509A3.49 3.49 0 0 0 4.51 2.021a3.492 3.492 0 0 0-3.488 3.488c0 1.086.239 1.905 1.018 2.546l.007 6.154c0 .435.403.734.833.734h5.117V5.509zM4.5 7C3.675 7 3 6.327 3 5.5C3 4.674 3.675 4 4.5 4C5.326 4 6 4.674 6 5.5S5.326 7 4.5 7zm11.812-1.968H9.006v5.938h7.306c.366 0 .662-.3.662-.667V5.698a.662.662 0 0 0-.662-.666zm-.296 2.999h-2.047V6.953h2.047v1.078z"/>`),
		g.Group(children),
	)
}