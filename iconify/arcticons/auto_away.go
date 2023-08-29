package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoAway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5v37h28.35v-8.317h8.65V5.5h-37Z" opacity=".995"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.089 27.266l-3.148-9.865l-3.269 9.865m1.089-3.329h4.238m9.67-3.221l-2.058 6.535l-1.937-6.535l-2.059 6.535l-1.937-6.535m19.35 6.553l-2.543-6.535m4.843 0L35.3 29.612a1.419 1.419 0 0 1-1.332.987h-.363m-2.497-5.809a2.422 2.422 0 1 1-4.843 0v-1.604a2.422 2.422 0 1 1 4.843 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.076 27.256a1.051 1.051 0 0 1-.968-.987V20.72M33.85 42.5l8.65-8.317"/>`),
		g.Group(children),
	)
}