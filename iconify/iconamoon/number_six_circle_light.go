package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><path fill="currentColor" d="M13.627 7.411a.75.75 0 0 0-1.254-.822l1.254.822Zm-4.78 4.557a.75.75 0 0 0 1.255.822l-1.254-.822ZM14.25 14A2.25 2.25 0 0 1 12 16.25v1.5A3.75 3.75 0 0 0 15.75 14h-1.5ZM12 16.25A2.25 2.25 0 0 1 9.75 14h-1.5A3.75 3.75 0 0 0 12 17.75v-1.5ZM9.75 14A2.25 2.25 0 0 1 12 11.75v-1.5A3.75 3.75 0 0 0 8.25 14h1.5ZM12 11.75A2.25 2.25 0 0 1 14.25 14h1.5A3.75 3.75 0 0 0 12 10.25v1.5Zm.373-5.161l-3.525 5.38l1.254.821l3.525-5.379l-1.254-.822Z"/></g>`),
		g.Group(children),
	)
}