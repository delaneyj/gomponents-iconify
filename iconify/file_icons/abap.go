package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M121.637 386.089L42.53 483.76l-19.71.547V12.176L437.766 0v7.305L380.815 66.86l-259.178 8.817v310.412zM489.179 73.625l-.946 386.024L160.555 512V86.949L489.18 73.625zM289.252 243.689c13.005 6.299 28.484 9.989 45.995 7.378c20.383-3.04 34.737-9.568 44.792-16.666l-43.36-124.255l-47.427 133.543zm151.04 163.38L395.53 278.795c-38.137 17.234-82.675 23.328-123.36 12.993l-53.023 149.303s21.65 38.662 115.985 18.558s105.16-52.58 105.16-52.58z"/>`),
		g.Group(children),
	)
}