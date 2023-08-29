package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSouthAfrica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m1.375 41.72L19.686 57.14a27.998 27.998 0 0 1-3.484-2.038l15.543-15.234h27.123a27.907 27.907 0 0 1-1.447 3.853H33.375zM5.221 23.825c.428-1.4.963-2.755 1.596-4.054L19.217 32l-12.4 12.229a27.877 27.877 0 0 1-1.597-4.056L13.505 32l-8.284-8.175m28.154-3.547h24.044a27.76 27.76 0 0 1 1.448 3.855H31.744L16.203 8.897a27.979 27.979 0 0 1 3.483-2.037l13.689 13.418"/>`),
		g.Group(children),
	)
}