package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeFbx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#008f90" d="M16.597 2L8.252 6.378l-.371 2.704l.014 18.365L16.419 30" opacity=".75"/><path fill="#00393a" d="M7.895 11.896v3.033l5.435 2.746l3.157 2.031l.043-7.817l-8.635.007z"/><path fill="#008f90" d="m16.53 11.889l-.043 7.817l6.904-2.663l.755-5.161l-7.616.007z"/><path fill="#004748" d="m8.225 6.378l-.371 2.691l8.702-.605L16.597 2L8.225 6.378z"/><path fill="#009b9d" d="M22.375 3.222L16.597 2l-.041 6.464l7.549-1.276l-1.73-3.966z"/><path fill="#002526" d="m8.252 6.378l-.371 2.704l.014 18.365l8.647-15.551l-8.29-5.518z" opacity=".7"/><path fill="#006c6e" d="m7.895 27.447l4.735 1.414L16.419 30l.123-18.104l-8.647 15.551z" opacity=".5"/>`),
		g.Group(children),
	)
}