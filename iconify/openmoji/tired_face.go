package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TiredFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FCEA2B" d="M36 13.2c-12.572 0-22.8 10.228-22.8 22.8c0 12.572 10.228 22.8 22.8 22.8c12.572 0 22.8-10.228 22.8-22.8c0-12.572-10.228-22.8-22.8-22.8z"/><g stroke="#000" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="36" r="23" fill="none"/><path fill="none" stroke-linecap="round" stroke-linejoin="round" d="M21.88 23.92c5.102-.061 7.273-1.882 8.383-3.346"/><path d="M46.24 47.56c0-2.592-2.867-7.121-10.25-6.93c-6.974.181-10.22 4.518-10.22 7.111s4.271-1.611 10.05-1.492c6.317.13 10.43 3.903 10.43 1.311z"/><path fill="none" stroke-linecap="round" stroke-linejoin="round" d="M23.16 28.47c5.215 1.438 5.603.91 8.204 1.207c1.068.122-2.03 2.67-7.282 4.397M50.12 23.92c-5.102-.061-7.273-1.882-8.383-3.346m7.103 7.896c-5.215 1.438-5.603.91-8.204 1.207c-1.068.122 2.03 2.67 7.282 4.397"/></g>`),
		g.Group(children),
	)
}