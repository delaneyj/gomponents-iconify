package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UniversalAccessAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 10a9 9 0 1 0-18.001.001A9 9 0 0 0 19 10zm-9-7.4c.83 0 1.5.67 1.5 1.5s-.67 1.51-1.5 1.51c-.82 0-1.5-.68-1.5-1.51s.68-1.5 1.5-1.5zM3.4 7.36c0-.65 6.6-.76 6.6-.76s6.6.11 6.6.76s-4.47 1.4-4.47 1.4s1.69 8.14 1.06 8.38c-.62.24-3.19-5.19-3.19-5.19s-2.56 5.43-3.18 5.19c-.63-.24 1.06-8.38 1.06-8.38S3.4 8.01 3.4 7.36z"/>`),
		g.Group(children),
	)
}