package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyrightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M17 3.34a10 10 0 1 1-14.995 8.984L2 12l.005-.324A10 10 0 0 1 17 3.34zm-2.34 5.659a4.016 4.016 0 0 0-5.543.23a3.993 3.993 0 0 0 0 5.542a4.016 4.016 0 0 0 5.543.23a1 1 0 0 0-1.32-1.502c-.81.711-2.035.66-2.783-.116a1.993 1.993 0 0 1 0-2.766a2.016 2.016 0 0 1 2.783-.116A1 1 0 0 0 14.66 9z"/></g>`),
		g.Group(children),
	)
}