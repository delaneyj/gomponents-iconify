package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ipad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.22 1.417H6.11c-.865 0-1.566.702-1.566 1.566v25.313c0 .865.7 1.565 1.566 1.565h19.11c.866 0 1.566-.7 1.566-1.564V2.984c0-.865-.7-1.567-1.565-1.567zM15.667 29.3a.626.626 0 1 1 0-1.25a.626.626 0 0 1 0 1.249zm8.71-2.445a.313.313 0 0 1-.313.312H7.27a.314.314 0 0 1-.313-.312V4.3c0-.173.14-.313.313-.313h16.792c.172 0 .312.14.312.313l.002 22.555z"/>`),
		g.Group(children),
	)
}