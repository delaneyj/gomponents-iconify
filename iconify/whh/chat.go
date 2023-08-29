package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 832v192q-7-4-20-12t-54-44t-82-84q-49 12-100 12q-73 0-134.5-15T405 844t-66.5-49t-39-50t-11.5-41h96q129 0 229.5-45T773 526t59-206q192 97 192 272q0 143-128 240zm-64-512zM384 640q-51 0-100-12q-41 48-80 83t-58 46l-18 11V576Q0 479 0 336q0-91 51.5-168.5T191.5 45t193-45t193 45T717 167.5T768 336t-52 161t-139 106.5T384 640zm0-576q-87 0-160.5 34T107 191T64 319.5T107 448t116.5 93.5T384 576t160.5-34.5T661 448t43-128.5T661 191T544.5 98T384 64z"/>`),
		g.Group(children),
	)
}