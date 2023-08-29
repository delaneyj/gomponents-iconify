package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Automobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M480 960q0-66-47-113t-113-47t-113 47t-47 113t47 113t113 47t113-47t47-113zm36-320h1016l-89-357q-2-8-14-17.5t-21-9.5H640q-9 0-21 9.5T605 283zm1372 320q0-66-47-113t-113-47t-113 47t-47 113t47 113t113 47t113-47t47-113zm160-96v384q0 14-9 23t-23 9h-96v128q0 80-56 136t-136 56t-136-56t-56-136v-128H512v128q0 80-56 136t-136 56t-136-56t-56-136v-128H32q-14 0-23-9t-9-23V864q0-93 65.5-158.5T224 640h28l105-419q23-94 104-157.5T640 0h768q98 0 179 63.5T1691 221l105 419h28q93 0 158.5 65.5T2048 864z"/>`),
		g.Group(children),
	)
}