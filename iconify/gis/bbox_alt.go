package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BboxAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M0 3v20h8.5v3h3v-3H20v-9h1.5v-3H20V3H0zm80 0v8h-1.5v3H80v9h8.5v3h3v-3h8.5V3H80zM5 8h10v10H5V8zm80 0h10v10H85V8zm-60.5 3v3h6v-3h-6zm9 0v3h6v-3h-6zm9 0v2.941L9.836 37.953a3.5 3.5 0 0 0-.06.047H8.5v1.979a3.5 3.5 0 0 0 0 1.587V44h1.063l1.417 3H8.5v6h3v-4.902l17.57 37.156a3.5 3.5 0 0 0 .479.746H24.5v3h6v-2.205a3.5 3.5 0 0 0 3 .225V89h6v-3h-4.535l55.494-50.994a3.5 3.5 0 0 0 .006-.006H91.5v-1.79a3.5 3.5 0 0 0 0-1.564V29h-2.586L48.5 13.719V11h-6zm9 0v3h6v-3h-6zm9 0v3h6v-3h-6zm9 0v3h6v-3h-6zm-23.682 9.19l35.723 13.505l-48.176 44.268l-17.058-36.078l29.511-21.696zM8.5 29v6h3v-6h-3zm80 9v6h3v-6h-3zm0 9v6h3v-6h-3zm-80 9v6h3v-6h-3zm80 0v6h3v-6h-3zm-80 9v6h3v-6h-3zm80 0v6h3v-6h-3zm-80 9v3H0v20h20v-8h1.5v-3H20v-9h-8.5v-3h-3zm80 0v3H80v9h-1.5v3H80v8h20V77h-8.5v-3h-3zM5 82h10v10H5V82zm80 0h10v10H85V82zm-42.5 4v3h6v-3h-6zm9 0v3h6v-3h-6zm9 0v3h6v-3h-6zm9 0v3h6v-3h-6z" color="currentColor"/>`),
		g.Group(children),
	)
}