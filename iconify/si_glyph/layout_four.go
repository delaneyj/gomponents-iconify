package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.956 6.326c0 .35-.299.632-.669.632H1.689c-.371 0-.67-.282-.67-.632V.633c0-.35.299-.633.67-.633h8.598c.37 0 .669.283.669.633v5.693zM17 15.333a.669.669 0 0 1-.671.667H7.714a.668.668 0 0 1-.671-.667V8.666c0-.368.3-.666.671-.666h8.615c.37 0 .671.298.671.666v6.667zm-11.024 0c0 .369-.3.667-.668.667H1.691a.668.668 0 0 1-.671-.667V8.666c0-.368.299-.666.671-.666h3.617c.368 0 .668.298.668.666v6.667zM17 6.345c0 .351-.324.634-.726.634h-3.528c-.398 0-.723-.283-.723-.634V.635c0-.352.324-.635.723-.635h3.528c.401 0 .726.283.726.635v5.71z"/>`),
		g.Group(children),
	)
}