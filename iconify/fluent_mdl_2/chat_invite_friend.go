package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatInviteFriend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1762 1590q66 32 118 80t90 108t58 128t20 142h-128q0-79-30-149t-83-122t-122-82t-149-31q-79 0-149 30t-122 83t-82 122t-31 149h-128q0-73 20-141t57-128t90-108t119-81q-75-54-116-135t-42-175q0-79 30-149t82-122t122-83t150-30q79 0 149 30t122 82t83 123t30 149q0 93-41 174t-117 136zm-226-54q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100q0 53 20 99t55 82t81 55t100 20zm512-1408v1024q0-48-9-84t-25-67t-40-60t-54-63V256H128v1152h256v267l267-267h390q8 34 21 66t31 62H704l-448 448v-448H0V128h2048z"/>`),
		g.Group(children),
	)
}