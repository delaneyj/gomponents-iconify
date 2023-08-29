package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartFillSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.655 14.916v-.001h-.002l-.006-.003l-.018-.01a22.066 22.066 0 0 1-3.744-2.584C2.045 10.731 0 8.35 0 5.5C0 2.836 2.086 1 4.25 1C5.797 1 7.153 1.802 8 3.02C8.847 1.802 10.203 1 11.75 1C13.914 1 16 2.836 16 5.5c0 2.85-2.044 5.231-3.886 6.818a22.094 22.094 0 0 1-3.433 2.414a7.152 7.152 0 0 1-.31.17l-.018.01l-.008.004a.75.75 0 0 1-.69 0Z"/>`),
		g.Group(children),
	)
}