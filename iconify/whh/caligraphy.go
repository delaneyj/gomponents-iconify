package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caligraphy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 704q-53 0-90.5 37.5T704 832q0 80-56 136t-136 56q-106 0-181-75t-75-181t75-181t181-75q53 0 90.5-37.5T640 384q0-80-65.5-136T416 192q-95 0-159.5 44.5T192 352q0 40-28 68t-68 28t-68-28t-28-68q0-104 52-185T191 43T384 0q104 0 192.5 51.5t140 140T768 384q0 106-75 181t-181 75q-53 0-90.5 47T384 800t37.5 113t90.5 47t90.5-37.5T640 832q0-80 56-136t136-56t136 56t56 136q-4-4-18.5-19T986 793.5t-18-17t-19-17.5l-18-15l-19.5-14l-18.5-10.5l-20-9l-19.5-4.5l-21.5-2z"/>`),
		g.Group(children),
	)
}