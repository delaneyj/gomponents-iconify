package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Doxygen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M365.469 69.597C301.552 5.018 222.01 0 199.154 0H0v511.996h117.932V117.932h81.222c22.247 0 57.2 9.069 82.495 34.625c25.708 25.974 38.474 65.218 37.944 116.64c-.532 51.408-23.237 85.093-71.456 106.01c-42.82 18.574-88.246 18.858-88.7 18.858v117.93c6.811.07 67.678-.28 131.472-26.83c93.516-38.917 145.582-115.184 146.609-214.75c1.08-104.72-38.585-167.006-72.05-200.818z"/>`),
		g.Group(children),
	)
}