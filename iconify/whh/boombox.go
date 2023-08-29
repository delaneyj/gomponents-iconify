package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boombox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 896H64q-26 0-45-19T0 832V320q0-26 18.5-45T64 256h64V64q0-26 19-45t45-19h640q26 0 45 19t19 45v192h64q26 0 45 19t19 45v512q0 26-19 45t-45 19zM768 160q0-13-9.5-22.5T736 128H288q-13 0-22.5 9.5T256 160v96h512v-96zm0 224q-80 0-136 56t-56 136q0 73 50 128H398q50-55 50-128q0-80-56-136t-136-56t-136 56t-56 136t56 136t136 56h512q80 0 136-56t56-136t-56-136t-136-56zm0 320q-53 0-90.5-37.5T640 576t37.5-90.5T768 448t90.5 37.5T896 576t-37.5 90.5T768 704zM128 576q0-53 37.5-90.5T256 448t90.5 37.5T384 576t-37.5 90.5T256 704t-90.5-37.5T128 576z"/>`),
		g.Group(children),
	)
}