package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M1.25 21a.75.75 0 0 1 .75-.75h20a.75.75 0 0 1 0 1.5H2a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd" opacity=".5"/><path fill-rule="evenodd" d="M3.002 13C3 12.688 3 12.355 3 12V7c0-1.886 0-2.828.586-3.414C4.172 3 5.114 3 7 3h6c1.886 0 2.828 0 3.414.586c.235.235.376.527.46.914H18c2.526 0 4.75 1.812 4.75 4.25c0 2.438-2.224 4.25-4.75 4.25H3.002Zm13.994-7c.004.3.004.632.004 1v4.5h1c1.892 0 3.25-1.322 3.25-2.75S19.892 6 18 6h-1.004Z" clip-rule="evenodd"/><path d="M9 18h2c2.829 0 4.243 0 5.122-.879c.768-.768.864-1.946.877-4.121H3.002c.012 2.175.109 3.353.877 4.121C4.758 18 6.172 18 9 18Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}