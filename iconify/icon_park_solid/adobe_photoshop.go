package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobePhotoshop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAdobePhotoshop0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M14 15v18"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M16 25c3 0 7-.8 7-5s-4-5-7-5h-2v10h2Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M34 21.025c-1 0-6-.498-6 2.987c0 3.486 6 2.49 6 5.976c0 3.485-6 2.987-6 2.987"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAdobePhotoshop0)"/>`),
		g.Group(children),
	)
}