package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headstone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M13.5 12a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5Zm.5 3a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4Zm-2 2.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5Zm4.5-.5a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3Z"/><path d="M31 31H1v-2a4.001 4.001 0 0 1 3.14-3.907A3.008 3.008 0 0 1 6 23.17V8a4 4 0 0 1 4-4h.803A5.996 5.996 0 0 1 16 1a5.996 5.996 0 0 1 5.197 3H22a4 4 0 0 1 4 4v15.17a3.008 3.008 0 0 1 1.86 1.923A4.001 4.001 0 0 1 31 29v2Zm-2-2a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2h26Zm-4-4H7a1 1 0 0 0-1 1h20a1 1 0 0 0-1-1Zm-1-1V8a2 2 0 0 0-2-2h-1.738a.545.545 0 0 1-.503-.37a4.002 4.002 0 0 0-7.518 0a.545.545 0 0 1-.503.37H10a2 2 0 0 0-2 2v16h1V8.5A1.5 1.5 0 0 1 10.5 7H13a3 3 0 1 1 6 0h2.5A1.5 1.5 0 0 1 23 8.5V24h1Zm-2 0V8.5a.5.5 0 0 0-.5-.5h-2.571A.929.929 0 0 1 18 7.071V7a2 2 0 1 0-4 0v.071a.929.929 0 0 1-.929.929H10.5a.5.5 0 0 0-.5.5V24h12Z"/></g>`),
		g.Group(children),
	)
}