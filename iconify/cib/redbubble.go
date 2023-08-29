package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redbubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.177 21.766h-4.266a.43.43 0 0 1-.427-.427V10.631a.43.43 0 0 1 .427-.427h3.953c2.969 0 3.594 1.75 3.594 3.214c0 .849-.224 1.521-.672 2.016c1.089.448 1.672 1.458 1.672 2.922c0 2.135-1.599 3.411-4.281 3.411zm-6.193 0H7.125a.43.43 0 0 1-.427-.427V10.631a.43.43 0 0 1 .427-.427h4.141c2.583 0 4.125 1.391 4.125 3.724c0 1.552-.776 2.771-2.036 3.266l2.948 3.859a.432.432 0 0 1-.318.714zM16 0C7.161 0 0 7.161 0 16s7.161 16 16 16s16-7.161 16-16S24.839 0 16 0z"/>`),
		g.Group(children),
	)
}