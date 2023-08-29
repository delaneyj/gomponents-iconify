package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M9 11.657V19c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C10.602 22 11.068 22 12 22c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C15 20.398 15 19.932 15 19v-7.343c0-.818 0-1.226.152-1.594c.153-.367.442-.657 1.02-1.235l3.242-3.242c.29-.29.434-.434.51-.618C20 4.785 20 4.58 20 4.172V4c0-.943 0-1.414-.293-1.707C19.414 2 18.943 2 18 2H6c-.943 0-1.414 0-1.707.293C4 2.586 4 3.057 4 4v.172c0 .408 0 .613.076.796c.076.184.22.329.51.618l3.242 3.242c.578.578.868.868 1.02 1.235c.152.368.152.776.152 1.594Z"/><path d="M15 10H9m3 3v2M4.5 5h15" opacity=".5"/></g>`),
		g.Group(children),
	)
}