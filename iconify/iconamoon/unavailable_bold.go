package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnavailableBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19.75A7.75 7.75 0 0 1 4.25 12h-2.5c0 5.66 4.59 10.25 10.25 10.25v-2.5Zm0-15.5A7.75 7.75 0 0 1 19.75 12h2.5c0-5.661-4.59-10.25-10.25-10.25v2.5ZM4.25 12c0-2.14.866-4.076 2.27-5.48L4.752 4.752A10.222 10.222 0 0 0 1.75 12h2.5Zm2.27-5.48A7.722 7.722 0 0 1 12 4.25v-2.5c-2.83 0-5.394 1.149-7.248 3.002L6.52 6.52Zm-1.768 0L17.48 19.248l1.768-1.768L6.52 4.752L4.752 6.52ZM19.75 12c0 2.14-.866 4.076-2.27 5.48l1.768 1.768A10.221 10.221 0 0 0 22.25 12h-2.5Zm-2.27 5.48A7.722 7.722 0 0 1 12 19.75v2.5c2.83 0 5.394-1.149 7.248-3.002L17.48 17.48Z"/>`),
		g.Group(children),
	)
}