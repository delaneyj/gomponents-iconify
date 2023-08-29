package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paypal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.47 3.544h8c1.639 0 2.945.775 3.626 1.971c1.22.872 1.847 2.4 1.512 4.138c-.521 2.712-3.183 4.91-5.944 4.91h-2l-1.134 5.892H6.398l.23-1.199h-3.18L6.47 3.544Zm1.622 1.964h6c1.657 0 2.746 1.32 2.433 2.946c-.313 1.627-1.91 2.946-3.566 2.946h-4l-1.134 5.892h-2L8.092 5.508Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}