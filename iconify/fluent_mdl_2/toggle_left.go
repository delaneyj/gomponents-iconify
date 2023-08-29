package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1024q0 106-40 199t-109 163t-163 110t-200 40H512q-106 0-199-40t-163-109t-110-163t-40-200q0-106 40-199t109-163t163-110t200-40h1024q106 0 199 40t163 109t110 163t40 200zm-512 384q79 0 149-30t122-82t83-122t30-150q0-79-30-149t-82-122t-123-83t-149-30H512q-80 0-149 30t-122 82t-83 123t-30 149q0 80 30 149t82 122t122 83t150 30h1024zM512 768q53 0 99 20t82 55t55 81t20 100q0 53-20 99t-55 82t-81 55t-100 20q-53 0-99-20t-82-55t-55-81t-20-100q0-53 20-99t55-82t81-55t100-20z"/>`),
		g.Group(children),
	)
}