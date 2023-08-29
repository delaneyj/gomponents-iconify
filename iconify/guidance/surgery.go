package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Surgery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M16 19.5H6.5v-1a3 3 0 0 1 3-3c1 0 14.4 1.5 14.4 1.5M0 22.5h24m-9.5-9l-2.918-2.085a1.39 1.39 0 0 1-.47-1.677l.985-2.297C12.466 7.226 13.87 6.5 16 6.5s3.534.726 3.903.94l.985 2.298a1.39 1.39 0 0 1-.47 1.678L17.5 13.5M.5.5l1.666 1.666M4.5 17.9s-1 1.6-2.25 1.6A1.75 1.75 0 0 1 .5 17.75c0-.966.783-1.746 1.75-1.746c1.25 0 2.25 1.596 2.25 1.596v.3ZM15.852 4.5s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75c.967 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25h-.3ZM4.221 1.544L6.5 2v.25L2.25 6.5H2l-.456-2.28a2.275 2.275 0 0 1 2.677-2.676Z"/>`),
		g.Group(children),
	)
}