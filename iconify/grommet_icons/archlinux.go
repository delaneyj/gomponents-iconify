package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archlinux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.995 0c-1.068 2.619-1.712 4.332-2.901 6.873c.729.773 1.624 1.673 3.077 2.69c-1.562-.643-2.628-1.289-3.425-1.959C7.224 10.78 4.84 15.304 0 24c3.804-2.196 6.752-3.55 9.5-4.066a6.964 6.964 0 0 1-.18-1.63l.004-.121c.06-2.437 1.328-4.311 2.83-4.184c1.501.127 2.668 2.207 2.608 4.644c-.011.459-.063.9-.153 1.309c2.717.532 5.634 1.882 9.387 4.048c-.74-1.362-1.4-2.59-2.031-3.76c-.994-.77-2.03-1.771-4.143-2.856c1.452.377 2.493.813 3.303 1.3C14.713 6.746 14.195 5.16 11.995 0Z"/>`),
		g.Group(children),
	)
}