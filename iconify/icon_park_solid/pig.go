package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPig0"><g fill="none"><path stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M14.054 9.644a9.115 9.115 0 0 1 1.414 1.845a15.95 15.95 0 0 1 8.483-2.426c3.146 0 6.08.906 8.555 2.471c.4-.691.886-1.337 1.44-1.89c2.521-2.516 6.946-3.624 8.991-1.583c2.045 2.04.934 6.456-1.587 8.972a9.358 9.358 0 0 1-2.638 1.824a15.89 15.89 0 0 1 1.24 6.175c0 8.819-7.164 15.968-16 15.968C15.113 41 7.95 33.85 7.95 25.032c0-2.204.447-4.304 1.256-6.214a9.34 9.34 0 0 1-2.556-1.785c-2.522-2.516-3.632-6.932-1.587-8.972c2.045-2.04 6.47-.933 8.99 1.583Z"/><ellipse cx="24" cy="29" fill="#fff" stroke="#fff" stroke-width="4" rx="8" ry="7"/><circle cx="17" cy="18" r="2" fill="#fff"/><circle cx="21" cy="29" r="2" fill="#000"/><circle cx="31" cy="18" r="2" fill="#fff"/><circle cx="27" cy="29" r="2" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPig0)"/>`),
		g.Group(children),
	)
}