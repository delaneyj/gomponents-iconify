package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 0C8.486 0 4 4.486 4 10v3H2.5A2.503 2.503 0 0 0 0 15.5v14C0 30.878 1.122 32 2.5 32h23c1.378 0 2.5-1.122 2.5-2.5v-14c0-1.378-1.122-2.5-2.5-2.5H24v-3c0-5.514-4.486-10-10-10zM5 10c0-4.962 4.038-9 9-9s9 4.038 9 9v3h-3v-3c0-3.309-2.691-6-6-6s-6 2.691-6 6v3H5v-3zm14 3H9v-3c0-2.757 2.243-5 5-5s5 2.243 5 5v3zm-6.707 8l-5 5H2.707l5-5h4.586zm6 0l-5 5H8.707l5-5h4.586zm6 0l-5 5h-4.586l5-5h4.586zM27 21v5h-6.293l5-5H27zM1.293 26H1v-5h5.293l-5 5zM25.5 31h-23c-.827 0-1.5-.673-1.5-1.5V27h26v2.5c0 .827-.673 1.5-1.5 1.5zM27 15.5V20H1v-4.5c0-.827.673-1.5 1.5-1.5h23c.827 0 1.5.673 1.5 1.5z"/>`),
		g.Group(children),
	)
}