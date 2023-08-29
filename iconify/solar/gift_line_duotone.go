package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M22 12H2M12 2v20" opacity=".5"/><path stroke-linecap="round" d="M13 12a4 4 0 0 0 4 4m-6-4a4 4 0 0 1-4 4"/><path d="M12 10.035a3.247 3.247 0 0 1 2.46-3.15c1.603-.4 3.056 1.052 2.655 2.656a3.247 3.247 0 0 1-3.15 2.46H12v-1.966Zm0 0a3.247 3.247 0 0 0-2.46-3.15c-1.603-.4-3.056 1.052-2.655 2.656a3.247 3.247 0 0 0 3.15 2.46H12v-1.966Z"/><path d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z"/></g>`),
		g.Group(children),
	)
}