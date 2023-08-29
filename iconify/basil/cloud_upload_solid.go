package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUploadSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.46 7.284a5.296 5.296 0 0 1 9.734 2.543a4.386 4.386 0 1 1 .17 8.77H7A5.75 5.75 0 1 1 8.46 7.284Zm6.626 5.185a.75.75 0 0 1-1.055.117L12.75 11.56V15a.75.75 0 0 1-1.5 0v-3.44l-1.282 1.025a.75.75 0 1 1-.937-1.172l2.498-1.997a.747.747 0 0 1 .465-.167h.008c.18 0 .344.064.473.17l2.494 1.994a.75.75 0 0 1 .117 1.055Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}