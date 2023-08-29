package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerAltEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<defs><path id="gisLayerAltEdit0" fill="currentColor" d="M90.361.872a2.977 2.977 0 0 0-4.209 0l-34.81 34.81c-.29.29-.508.626-.653.985L45.934 53.29c-.218.687.117.98.858.753l16.541-4.732c.359-.145.695-.362.985-.653l34.81-34.81a2.977 2.977 0 0 0 0-4.21zm-7.576 11.786l4.557 4.557l-25.128 25.13l-4.558-4.559z" color="currentColor"/></defs><path fill="currentColor" d="M50.049 24.035c-1.27-.007-2.493.267-3.39.76L1.382 49.64c-1.847 1.012-1.847 2.655 0 3.668l45.275 24.845c1.846 1.013 4.838 1.013 6.684 0L98.617 53.31c1.847-1.013 1.847-2.656 0-3.668L79.266 39.02l-4.43 4.903l13.76 7.55L50 72.655l-38.594-21.18L50 30.295l.78.429l4.49-4.87l-1.928-1.058c-.874-.48-2.057-.753-3.293-.76ZM4.727 60.857l-3.344 1.834c-1.847 1.013-1.847 2.656 0 3.668l45.275 24.846c1.846 1.013 4.838 1.013 6.684 0L98.617 66.36c1.846-1.013 1.845-2.655-.002-3.668l-3.342-1.834l-6.683 3.666l.004.002L50 85.705l-38.596-21.18l.004-.002z" color="currentColor"/><use href="#gisLayerAltEdit0" color="currentColor"/><use href="#gisLayerAltEdit0" color="currentColor"/>`),
		g.Group(children),
	)
}