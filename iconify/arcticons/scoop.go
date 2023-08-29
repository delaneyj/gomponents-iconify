package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.064 21.951a8.936 8.936 0 0 1 17.872 0m0 6.399a8.936 8.936 0 1 1-17.872 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="4.75" d="M15.064 28.35v-6.4m17.872.001v6.398m-21.358-9.306h3.701m-3.701 6.144h3.289m-3.289 6.145h3.701M32.72 19.043h3.702m-3.339 6.144h3.339m-3.702 6.145h3.702M20.207 13.634l-2.92-2.92m10.483 2.954l2.95-2.949m-9.728 11.379h6.065m-6.065 6.179h6.065"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}