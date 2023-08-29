package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weightscale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M971.23 904q-3 50-43 84.5t-93 34.5h-640q-53 0-93-34.5t-44-84.5l-57-748q-4-64 36-110t102-46h747q62 0 102.5 46t36.5 110zm-458-776q-100 0-186.5 49.5T193.23 320l218 87l-81-161q-9-10-9-23t9.5-22.5t22.5-9.5t22 10l161 184q46 4 76 23l221-88q-47-93-133.5-142.5T513.23 128z"/>`),
		g.Group(children),
	)
}