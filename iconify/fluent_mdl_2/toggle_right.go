package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1898 662q73 73 111 166t39 196q0 103-38 196t-112 166q-73 73-166 111t-196 39H512q-103 0-196-38t-166-112q-73-73-111-166T0 1024q0-103 38-196t112-166q73-73 166-111t196-39h1024q103 0 196 38t166 112zm-362 746q79 0 149-30t122-82t83-122t30-150q0-79-30-149t-82-122t-123-83t-149-30H512q-80 0-149 30t-122 82t-83 123t-30 149q0 80 30 149t82 122t122 83t150 30h1024zm0-640q53 0 99 20t82 55t55 81t20 100q0 53-20 99t-55 82t-81 55t-100 20q-53 0-99-20t-82-55t-55-81t-20-100q0-53 20-99t55-82t81-55t100-20z"/>`),
		g.Group(children),
	)
}