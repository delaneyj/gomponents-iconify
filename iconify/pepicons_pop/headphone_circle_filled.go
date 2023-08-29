package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopHeadphoneCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M20.88 18.069a1 1 0 0 1-1.898-.63a10.4 10.4 0 0 0 .518-3.273c0-4.456-2.756-7.412-6.5-7.412S6.5 9.71 6.5 14.166c0 1.14.178 2.247.518 3.273a1 1 0 0 1-1.898.63a12.4 12.4 0 0 1-.62-3.903c0-5.53 3.619-9.412 8.5-9.412s8.5 3.882 8.5 9.412c0 1.354-.212 2.673-.62 3.902Z"/><path d="M8.977 20.034a3 3 0 0 1-2.942-3.04v-.022a2.978 2.978 0 0 1 3.018-2.938h.017a1 1 0 0 1 .98 1.014l-.054 4a1 1 0 0 1-1.019.986ZM17.089 14a3 3 0 0 1 2.942 3.04v.022A2.978 2.978 0 0 1 17.013 20h-.016a1 1 0 0 1-.981-1.014l.054-4A1 1 0 0 1 17.089 14Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopHeadphoneCircleFilled0)"/></g>`),
		g.Group(children),
	)
}