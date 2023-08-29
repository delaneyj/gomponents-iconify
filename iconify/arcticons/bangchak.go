package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bangchak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.616 10.407c-16.55.22-27.47 13.428-19.948 21.34c7.132 7.501 22.07.2 19.948-21.34a84.741 84.741 0 0 1 4.56-.117a69.263 69.263 0 0 0-.56-4.591c-23.452-2.44-44.305 18.074-28.51 32.568c11.327 10.393 31.678 1.865 29.07-27.977"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.484 18.34c2.075-3.903 6.724-11.323 23.131-12.64M17.47 10.907c-3.218 3.98-3.617 11.16-3.217 15.368m-4.941 8.682c7.464 3.485 15.664.64 18.289-1.373"/>`),
		g.Group(children),
	)
}