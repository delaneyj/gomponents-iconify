package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yaml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="m.5 5.629l21.754 34.15v21.646h13.959V39.779l22.775-34.15h-15.02L30.021 27.617L16.189 5.629H.5zm63.994.13l-23.66 55.798h11.189l5.139-12.408h25.266l4.252 12.408h11.957L75.937 5.76H64.496zm5.992 11.774l7.744 20.475H61.843l8.643-20.475zm16.195 50.139v54.45H127.5v-11.635H98.636V67.672H86.681zm-64.428.011v54.687h11.734V84.647l12.28 25.355H55.5l12.7-26.246v38.602h11.256V67.682h-15.37L50.45 92.414L37.464 67.682H22.253z"/>`),
		g.Group(children),
	)
}