package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Projectfork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 256q-64 0-103-52q-181 40-309.5 202T199 791q26 17 41.5 45t15.5 60q0 53-37.5 90.5T128 1024t-90.5-37.5T0 896q0-35 17.5-64T64 786V64q0-26 18.5-45t45-19T173 19t19 45v355q80-136 199-225.5T651 76q15-34 46.5-55T768 0q53 0 90.5 37.5T896 128t-37.5 90.5T768 256zM128.5 832Q102 832 83 850.5t-19 45T83 941t45.5 19t45-19t18.5-45.5t-18.5-45t-45-18.5zm640-768Q742 64 723 83t-19 45.5t19 45t45.5 18.5t45-19t18.5-45t-18.5-45t-45-19z"/>`),
		g.Group(children),
	)
}