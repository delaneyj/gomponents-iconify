package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="m14 8.79l7.314-1.27a1.5 1.5 0 0 1 .242-.02c.801 0 1.444.664 1.444 1.475v9.786c0 .72-.511 1.34-1.213 1.456l-7.705 1.276a.499.499 0 0 1-.18-.002l-7.647-1.267A1.5 1.5 0 0 1 5 18.744V9.011a1.5 1.5 0 0 1 1.756-1.478L14 8.79Z" opacity=".2"/><path fill-rule="evenodd" d="M13.08 7.304L5.244 6.019A1.5 1.5 0 0 0 3.5 7.5v9.738a1.5 1.5 0 0 0 1.268 1.482l8.155 1.275a.5.5 0 0 0 .577-.494V7.797a.5.5 0 0 0-.42-.493Zm-8-.298l7.42 1.216v10.694L4.923 17.73a.5.5 0 0 1-.423-.493V7.5a.5.5 0 0 1 .58-.494Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M21 6a1.5 1.5 0 0 0-.243.02L12.92 7.303a.5.5 0 0 0-.419.493V19.5a.5.5 0 0 0 .577.494l8.155-1.276a1.5 1.5 0 0 0 1.268-1.481V7.5A1.5 1.5 0 0 0 21 6Zm.077 11.73L13.5 18.916V8.222l7.42-1.216a.501.501 0 0 1 .58.494v9.737a.5.5 0 0 1-.423.494Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}