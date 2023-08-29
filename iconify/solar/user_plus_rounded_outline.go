package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserPlusRoundedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.25 6a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0ZM12 2.75a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5Z" clip-rule="evenodd"/><path d="M17.75 16.667a.75.75 0 1 0-1.5 0v.583h-.584a.75.75 0 0 0 0 1.5h.584v.583a.75.75 0 0 0 1.5 0v-.583h.583a.75.75 0 0 0 0-1.5h-.583v-.583Z"/><path fill-rule="evenodd" d="M16.812 13.254a4.75 4.75 0 1 1-2.901 8.355A12.87 12.87 0 0 1 12 21.75c-2.04 0-3.922-.47-5.322-1.27C5.3 19.692 4.25 18.49 4.25 17s1.05-2.692 2.428-3.48c1.4-.8 3.283-1.27 5.322-1.27c1.805 0 3.482.368 4.812 1.004ZM17 14.75a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5Zm-4.2 5.472A4.729 4.729 0 0 1 12.25 18c0-1.671.863-3.14 2.168-3.987A10.956 10.956 0 0 0 12 13.75c-1.827 0-3.444.425-4.578 1.073C6.267 15.483 5.75 16.28 5.75 17s.517 1.517 1.672 2.177c1.134.648 2.751 1.073 4.578 1.073c.271 0 .539-.01.8-.028Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}