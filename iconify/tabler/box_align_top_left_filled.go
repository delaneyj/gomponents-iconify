package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxAlignTopLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M10 3H5a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm5 0a1 1 0 0 1 .117 1.993L14.99 5a1 1 0 0 1-.117-1.993L15 3zm5 0a1 1 0 0 1 .117 1.993L19.99 5a1 1 0 0 1-.117-1.993L20 3zm0 5a1 1 0 0 1 .117 1.993L19.99 10a1 1 0 0 1-.117-1.993L20 8zm0 6a1 1 0 0 1 .117 1.993L19.99 16a1 1 0 0 1-.117-1.993L20 14zM4 14a1 1 0 0 1 .117 1.993L3.99 16a1 1 0 0 1-.117-1.993L4 14zm16 5a1 1 0 0 1 .117 1.993L19.99 21a1 1 0 0 1-.117-1.993L20 19zm-5 0a1 1 0 0 1 .117 1.993L14.99 21a1 1 0 0 1-.117-1.993L15 19zm-6 0a1 1 0 0 1 .117 1.993L8.99 21a1 1 0 0 1-.117-1.993L9 19zm-5 0a1 1 0 0 1 .117 1.993L3.99 21a1 1 0 0 1-.117-1.993L4 19z"/></g>`),
		g.Group(children),
	)
}