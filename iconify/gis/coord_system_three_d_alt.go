package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoordSystemThreeDAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M42.34 0v3.553h9.498l-9.78 11.828v2.844h15.883v-3.551H47.883l9.777-11.83V0H42.34zm5.16 22.5v32.158L16.748 73.404l2.604 4.27L50 58.99l30.648 18.684l2.604-4.27L52.5 54.658V22.5h-5zM.244 81.775l6.08 8.922L0 100h4.92l4.225-6.227l4.26 6.227h4.894l-6.322-9.303l6.078-8.922h-4.932l-3.978 5.871l-3.993-5.87H.244zm81.164 0l6.946 10.547V100h4.7v-7.678L100 81.775h-5.152l-4.15 6.495l-4.151-6.495h-5.139z" color="currentColor"/>`),
		g.Group(children),
	)
}