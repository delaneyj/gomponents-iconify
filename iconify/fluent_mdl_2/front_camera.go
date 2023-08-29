package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrontCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1258 1198q91 36 166 96t128 138t83 169t29 191h-128q0-106-40-199t-110-162t-163-110t-199-41q-106 0-199 40t-162 110t-110 163t-41 199H384q0-99 29-190t82-169t128-138t167-97q-71-54-110-133t-40-169q0-79 30-149t82-122t122-83t150-30q79 0 149 30t122 82t83 123t30 149q0 90-39 169t-111 133zM768 896q0 53 20 99t55 82t81 55t100 20q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100zm1280-640v1280h-267q-11-33-25-65t-31-63h195V384H128v1024h195q-17 31-31 63t-25 65H0V256h2048z"/>`),
		g.Group(children),
	)
}