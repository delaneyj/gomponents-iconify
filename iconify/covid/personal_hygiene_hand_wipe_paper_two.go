package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalHygieneHandWipePaperTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M23.083 7.916a3.583 3.583 0 0 1-7.166 0C15.917 5.229 19.5.75 19.5.75s3.583 4.479 3.583 7.166Zm-5.264 6.335c.859 2.035 2.006 3.466-.084 6h-4.818m-6-10.201c-3.314 0-6-1.074-6-2.4v13.2c0 1.326 2.686 2.4 6 2.4s6-1.074 6-2.4V7.65c0 1.326-2.687 2.4-6 2.4Z"/><path d="M6.917 10.05c3.314 0 6-1.075 6-2.4s-2.686-2.4-6-2.4s-6 1.075-6 2.4s2.686 2.4 6 2.4Zm-.6-2.4h1.2m-.6-2.4h6.47"/></g>`),
		g.Group(children),
	)
}