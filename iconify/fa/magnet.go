package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1536 704v128q0 201-98.5 362t-274 251.5T768 1536t-395.5-90.5t-274-251.5T0 832V704q0-26 19-45t45-19h384q26 0 45 19t19 45v128q0 52 23.5 90t53.5 57t71 30t64 13t44 2t44-2t64-13t71-30t53.5-57t23.5-90V704q0-26 19-45t45-19h384q26 0 45 19t19 45zM512 64v384q0 26-19 45t-45 19H64q-26 0-45-19T0 448V64q0-26 19-45T64 0h384q26 0 45 19t19 45zm1024 0v384q0 26-19 45t-45 19h-384q-26 0-45-19t-19-45V64q0-26 19-45t45-19h384q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}