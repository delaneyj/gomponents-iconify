package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powerjack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 832H128q-53 0-90.5-37.5T0 704V128q0-53 37.5-90.5T128 0h704q53 0 90.5 37.5T960 128v576q0 53-37.5 90.5T832 832zm0-512q0-13-9-22L661 136q-10-10-25-8H323q-15-2-25 8L136 298q-10 10-8 25v349q0 13 9.5 22.5T160 704h640q13 0 22.5-9.5T832 672V320zM672 576q-13 0-22.5-9.5T640 544v-64q0-13 9.5-22.5T672 448t22.5 9.5T704 480v64q0 13-9.5 22.5T672 576zM480 384q-13 0-22.5-9.5T448 352v-64q0-13 9.5-22.5T480 256t22.5 9.5T512 288v64q0 13-9.5 22.5T480 384zM288 576q-13 0-22.5-9.5T256 544v-64q0-13 9.5-22.5T288 448t22.5 9.5T320 480v64q0 13-9.5 22.5T288 576z"/>`),
		g.Group(children),
	)
}