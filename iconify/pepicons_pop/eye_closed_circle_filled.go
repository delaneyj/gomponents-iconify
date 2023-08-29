package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeClosedCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopEyeClosedCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M4.688 11.306a1 1 0 1 1 1.624-1.167c.304.423.773.827 1.38 1.177c1.293.746 3.085 1.184 5.012 1.184c1.928 0 3.72-.438 5.013-1.184c.607-.35 1.076-.754 1.38-1.177a1 1 0 1 1 1.624 1.167c-.48.669-1.164 1.257-2.004 1.742c-1.612.93-3.748 1.452-6.013 1.452c-2.264 0-4.4-.522-6.012-1.452c-.84-.485-1.524-1.073-2.004-1.742Z"/><path d="M14 14a1 1 0 1 0-2 0v2.5a1 1 0 1 0 2 0V14Zm-5.47-.242a1 1 0 0 1 1.94.485l-.5 2a1 1 0 1 1-1.94-.485l.5-2Zm8.44 0a1 1 0 0 0-1.94.485l.5 2a1 1 0 1 0 1.94-.485l-.5-2Zm2.555-2.465a1 1 0 1 0-1.414 1.414l2 2a1 1 0 0 0 1.414-1.414l-2-2Zm-13.447.034a1 1 0 0 1 1.48 1.346l-1.818 2a1 1 0 1 1-1.48-1.346l1.818-2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopEyeClosedCircleFilled0)"/></g>`),
		g.Group(children),
	)
}