package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenPadlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 43q35 0 60 25t25 60q0 21 22 21q21 0 21-21q0-52-38-90T192 0t-90 38t-38 90v64q-27 0-45.5 18.5T0 256v192q0 27 18.5 45.5T64 512h256q27 0 45.5-18.5T384 448V256q0-27-18.5-45.5T320 192H107v-64q0-35 25-60t60-25zm128 192q21 0 21 21v192q0 21-21 21H64q-21 0-21-21V256q0-21 21-21h256zM171 380v25q0 22 21 22t21-22v-25q43-15 43-60q0-27-18.5-45.5T192 256t-45.5 18.5T128 320q0 20 12 36.5t31 23.5zm21-81q21 0 21 21t-21 21t-21-21t21-21z"/>`),
		g.Group(children),
	)
}