package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Impresscms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 320q-40 0-68 18.5T832 384v384H704V416q0-40-28-68t-68-28t-68 28t-28 68v352H384V256l64-64h64v50q55-50 128-50q49 0 91.5 23t68.5 63q26-40 68.5-63t91.5-23q32 0 64 11v135q-34-18-96-18zm-736-64l-64 64v-72q-57-13-92.5-46T0 128q0-53 56-90.5T192 0t136 37.5t56 90.5t-56.5 90.5T192 256zm64 64v448H128V384l64-64h64z"/>`),
		g.Group(children),
	)
}