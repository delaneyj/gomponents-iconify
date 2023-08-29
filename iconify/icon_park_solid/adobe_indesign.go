package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeIndesign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSAdobeIndesign0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M16 15v18m16-16v16"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M28 33c-2.5 0-4-1.6-4-4s1.5-4 4-4h4v8h-4Z" clip-rule="evenodd"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSAdobeIndesign0)"/>`),
		g.Group(children),
	)
}