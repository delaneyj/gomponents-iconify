package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.986 17c0-.105-.004-.211-.038-.316l-2-6a.996.996 0 0 0-.56-.594a2.995 2.995 0 0 0-.269-3.914L12 .055L5.879 6.176a2.998 2.998 0 0 0-.27 3.914a.987.987 0 0 0-.559.594l-2 6a1.007 1.007 0 0 0-.038.316C3 17 3 22 3 22a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1s0-5-.014-5zM7.293 7.59L12 2.883l4.707 4.707a.999.999 0 0 1 0 1.414a1.025 1.025 0 0 1-1.414 0L13 6.711V12.5a1 1 0 0 1-2 0V6.711L8.707 9.004a1.025 1.025 0 0 1-1.414 0a.999.999 0 0 1 0-1.414zM6.721 12H9v.5c0 1.654 1.346 3 3 3s3-1.346 3-3V12h2.279l1.666 5H5.053l1.668-5zM5 21v-3h14v3H5z"/>`),
		g.Group(children),
	)
}