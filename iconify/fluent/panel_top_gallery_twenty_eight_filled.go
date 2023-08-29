package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelTopGalleryTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 4H5.754c-.888 0-1.703.308-2.346.824a2.01 2.01 0 0 0-.587.59a3.734 3.734 0 0 0-.817 2.336v12.5A3.75 3.75 0 0 0 5.754 24H22.25A3.75 3.75 0 0 0 26 20.25V7.75c0-.898-.315-1.721-.84-2.367a2.01 2.01 0 0 0-.544-.543A3.734 3.734 0 0 0 22.25 4H19v8h5.5v8.25a2.25 2.25 0 0 1-2.25 2.25H5.755a2.25 2.25 0 0 1-2.25-2.25V12H9V4Zm1.5 8h7V4h-7v8Z"/>`),
		g.Group(children),
	)
}