package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Symphony(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024q-53 0-90.5-37.5T768 896t37.5-90.5T896 768t90.5 37.5T1024 896t-37.5 90.5T896 1024zM288 448q5 2 16 6q33 14 49.5 20.5T399 494t43.5 22t37 23.5t34 28t25.5 32t20.5 38.5t11.5 45t5 53q0 124-83.5 206T288 1024q-59 0-103-11t-77.5-35.5T50 923T0 848l96-80q27 54 76 91t116 37q66 0 113-47t47-113q0-26-8-47t-18.5-34.5t-34.5-29t-43-24.5t-56-25q-5-2-16-6q-33-14-49.5-20.5T177 530t-43.5-22t-37-23.5t-34-28t-25.5-32T16.5 386T5 341t-5-53Q0 169 84.5 84.5T288 0q111 0 192.5 73.5T574 256H445q-11-55-55-91.5T288 128q-66 0-113 47t-47 113q0 25 12.5 47.5t37 42.5t49.5 35.5t61 34.5z"/>`),
		g.Group(children),
	)
}