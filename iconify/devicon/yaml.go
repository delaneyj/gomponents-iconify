package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yaml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path d="M22.254 39.778L.5 5.629h15.69l13.833 21.989L43.968 5.629h15.02L36.214 39.778v21.65h-13.96z"/><path fill="#cb171e" d="M82.428 49.149H57.162l-5.139 12.408H40.835l23.66-55.798h11.443l22.7 55.798H86.68l-4.25-12.408zm-4.197-11.14l-7.745-20.476l-8.642 20.476z"/><path d="M22.254 67.686v54.688h11.733V84.65l12.28 25.356h9.236l12.7-26.246v38.601h11.256V67.686H64.09l-13.638 24.73l-12.988-24.73zm105.248 42.804H98.639V67.67H86.682v54.455h40.82z"/>`),
		g.Group(children),
	)
}