package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOnLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M15 22v-4.343c0-.818 0-1.226.152-1.594c.153-.367.442-.656 1.02-1.235l3.242-3.242c.29-.29.434-.434.51-.617c.076-.184.076-.389.076-.797V10c0-.943 0-1.414-.293-1.707C19.414 8 18.943 8 18 8H6c-.943 0-1.414 0-1.707.293C4 8.586 4 9.057 4 10v.172c0 .408 0 .613.076.796c.076.184.22.329.51.618l3.242 3.242c.578.579.868.867 1.02 1.235c.152.368.152.776.152 1.594V22m6-6H9m-4.5-5h15M12 5V2M8 5L6 3m10 2l2-2m-6 16v2"/>`),
		g.Group(children),
	)
}