package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KartaKrakowska(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31.569" height="39" x="8.215" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.468 30.709v-6.396h2.094a2.148 2.148 0 0 1 0 4.296h-2.094m2.094 0l2.093 2.098m-9.298-6.394v6.396m3.438 0l-2.633-3.198l2.633-3.176m-2.633 3.176h-.805m10.909 3.179l2.078-6.377m2.079 6.396l-2.079-6.396m1.384 4.256h-2.771m-6.896 2.14v6.395m3.437 0l-2.633-3.197l2.651-3.198m-2.651 3.198h-.804m16.577-3.198l-1.599 6.395l-1.599-6.395l-1.599 6.395l-1.599-6.395m-5.656 4.277a2.119 2.119 0 0 0 4.237 0v-2.159a2.119 2.119 0 1 0-4.237 0Zm-2.696 2.118V43.5m3.438 0l-2.634-3.198l2.651-3.198m-2.651 3.198h-.804M14.115 42.8a1.789 1.789 0 0 0 1.568.7h.947a1.597 1.597 0 0 0 1.595-1.599h0a1.597 1.597 0 0 0-1.595-1.599h-1.046a1.597 1.597 0 0 1-1.595-1.599h0a1.597 1.597 0 0 1 1.595-1.599h.947a1.789 1.789 0 0 1 1.568.701m6.787 5.676l2.078-6.377m2.079 6.396l-2.079-6.396m1.384 4.257h-2.771m-12.22-4.257h26.428m-26.428-6.395h26.428m-26.428-6.396h26.428"/>`),
		g.Group(children),
	)
}