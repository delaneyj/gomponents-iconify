package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArcticonsPleading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m24 10.84l-11.4 6.58v13.16L24 37.16l11.4-6.58V17.42L24 10.84Z"/><circle cx="8.07" cy="14.8" r="3.11"/><circle cx="24" cy="5.61" r="3.11"/><circle cx="39.93" cy="14.8" r="3.11"/><circle cx="39.93" cy="33.2" r="3.11"/><circle cx="24" cy="42.39" r="3.11"/><circle cx="8.07" cy="33.2" r="3.11"/><path d="m12.6 30.58l-1.84 1.06M24 37.16v2.12m0-28.44V8.71m11.4 21.87l1.84 1.06M35.4 17.42l1.84-1.06M12.6 17.42l-1.84-1.06"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.504 29.331c.744-.234 1.527-.339 2.496-.339s1.753.105 2.497.339m-3.363-7.061c0-2.417-.12-5.118-2.447-5.118c-4.004 0-5.687 4.322-5.687 5.895c0 1.925 1.773 2.856 4.19 2.856s3.944-.936 3.944-3.633Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.915 20.175a6.007 6.007 0 0 0 5.217 3.024a6 6 0 0 0 1.974-.332"/><path fill="currentColor" d="M18.712 18.627a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Zm8.123 0a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.866 22.27c0-2.417.12-5.118 2.447-5.118c4.004 0 5.687 4.322 5.687 5.895c0 1.925-1.773 2.856-4.19 2.856s-3.944-.936-3.944-3.633Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.085 20.175a6.007 6.007 0 0 1-5.217 3.024a6.003 6.003 0 0 1-1.974-.332"/>`),
		g.Group(children),
	)
}