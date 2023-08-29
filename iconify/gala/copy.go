package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaCopy0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><rect id="galaCopy1" width="144" height="144" x="16" y="16" ry="32"/><path id="galaCopy2" d="M 95.999712,127.9996 A 31.999891,31.999891 0 0 1 127.9996,95.999712"/><path id="galaCopy3" d="m -239.99922,127.9996 a 31.999891,31.999891 0 0 1 31.99989,-31.999888" transform="scale(-1 1)"/><path id="galaCopy4" d="m -239.99922,-207.99933 a 31.999891,31.999891 0 0 1 31.99989,-31.99989" transform="scale(-1)"/><path id="galaCopy5" d="M 95.999712,-207.99933 A 31.999891,31.999891 0 0 1 127.9996,-239.99922" transform="scale(1 -1)"/><path id="galaCopy6" stroke-opacity="1" d="m 159.99949,239.99923 h 15.99995"/><path id="galaCopy7" stroke-opacity="1" d="m 159.99949,95.9997 h 15.99995"/><path id="galaCopy8" stroke-opacity="1" d="m 95.999709,159.99949 v 15.99995"/><path id="galaCopy9" stroke-opacity="1" d="m 239.99923,159.99949 v 15.99995"/></g>`),
		g.Group(children),
	)
}