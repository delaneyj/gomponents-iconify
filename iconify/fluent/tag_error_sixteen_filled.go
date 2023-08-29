package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagErrorSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.536 6.327L6.296 1.6a2.02 2.02 0 0 1 1.415-.585L10.975 1a2.007 2.007 0 0 1 2.021 2.014l-.019 2.574a5.5 5.5 0 0 0-7.428 7.31a2.008 2.008 0 0 1-.703-.454l-3.31-3.288a1.99 1.99 0 0 1 0-2.829Zm7.728-1.669a.774.774 0 0 0 1.09 0c.3-.298.3-.783 0-1.082a.774.774 0 0 0-1.09 0c-.3.3-.3.784 0 1.082ZM15 10.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0ZM10.5 8a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 1 0v-2a.5.5 0 0 0-.5-.5Zm0 5.125a.625.625 0 1 0 0-1.25a.625.625 0 0 0 0 1.25Z"/>`),
		g.Group(children),
	)
}