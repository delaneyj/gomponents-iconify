package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 11.657V19c0 .932 0 1.398.153 1.765a2 2 0 0 0 1.082 1.083C10.602 22 11.069 22 12 22c.932 0 1.398 0 1.766-.152a2 2 0 0 0 1.082-1.083C15 20.398 15 19.932 15 19v-7.343c0-.818 0-1.226.153-1.594L15.18 10H8.82c.01.02.02.042.028.063c.152.368.152.776.152 1.594Z" opacity=".5"/><path fill-rule="evenodd" d="M12 12.25a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/><path d="M20 4.172V4c0-.943 0-1.414-.293-1.707C19.414 2 18.943 2 18 2H6c-.943 0-1.414 0-1.707.293C4 2.586 4 3.057 4 4v.172c0 .408 0 .613.076.796L4.09 5h15.82c.005-.01.01-.02.014-.032C20 4.785 20 4.58 20 4.172Z" opacity=".5"/><path d="m16.171 8.828l3.243-3.242c.272-.273.416-.417.496-.586H4.09c.08.17.223.313.496.586l3.242 3.242c.545.545.833.833.992 1.172h6.36c.159-.339.447-.627.991-1.172Z"/></g>`),
		g.Group(children),
	)
}