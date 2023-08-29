package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m80 495.918l352.039.08h.006a24 24 0 0 0 24-24.008L456 354.472a23.9 23.9 0 0 0-4.239-13.613L344 184.511V121h40a24.028 24.028 0 0 0 24-24V16H104v81a24.027 24.027 0 0 0 24 24h39.917v64.621L60.276 340.834A23.9 23.9 0 0 0 56 354.509v117.409a24.029 24.029 0 0 0 24 24Zm119.917-300.287V89H136V48h240v41h-64v105.47L376.465 288H135.859ZM88 357.011L113.667 320H398.52L424 356.971L424.037 464L88 463.92Z"/>`),
		g.Group(children),
	)
}