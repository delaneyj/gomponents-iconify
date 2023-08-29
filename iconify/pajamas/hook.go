package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 11.125a3.875 3.875 0 0 0 7 2.292a3.875 3.875 0 0 0 7-2.292V7.002l-1.28 1.28l-1.49 1.488a.75.75 0 0 0 1.061 1.061l.208-.208v.502a2.375 2.375 0 1 1-4.75 0v-5.24a2.501 2.501 0 1 0-1.5 0v5.24a2.375 2.375 0 1 1-4.75 0v-.502l.208.208a.75.75 0 1 0 1.06-1.06L2.28 8.281L1 7.002v4.123ZM9 3.5a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}