package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gbx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#1666af"/><path fill="#fff" fill-rule="nonzero" d="M15.699 11.006v2.06h-4.7c-.61 0-1.361.214-1.92.61c-.693.492-1.08 1.229-1.08 2.327s.387 1.835 1.08 2.327c.559.396 1.31.61 1.92.61h2V21h-2a5.448 5.448 0 0 1-3.054-.973C6.717 19.155 6 17.789 6 16.003s.717-3.152 1.945-4.024a5.447 5.447 0 0 1 3.053-.973zm-5 6.182v-2.06h5V21H13.7v-3.812zM19 13.06V11h7v10h-9.002v-7.933h2v5.872H24v-5.878z"/></g>`),
		g.Group(children),
	)
}