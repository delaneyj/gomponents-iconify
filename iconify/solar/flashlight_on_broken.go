package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOnBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M15 22v-4.343c0-.818 0-1.226.152-1.594c.153-.367.442-.656 1.02-1.235l3.242-3.242c.29-.29.434-.434.51-.617c.076-.184.076-.389.076-.797V10c0-.943 0-1.414-.293-1.707C19.414 8 18.943 8 18 8h-1M9 22v-4.343c0-.818 0-1.226-.152-1.594c-.152-.367-.442-.656-1.02-1.235l-3.242-3.242c-.29-.29-.434-.434-.51-.617C4 10.785 4 10.58 4 10.171V10c0-.943 0-1.414.293-1.707C4.586 8 5.057 8 6 8h7m2 8H9m-4.5-5h15M12 5V2M8 5L6 3m10 2l2-2m-6 16v2"/>`),
		g.Group(children),
	)
}