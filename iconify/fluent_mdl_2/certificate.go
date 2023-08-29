package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Certificate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 1280q0 82-33 156t-95 130v458l-256-128l-256 128v-458q-61-55-94-129t-34-157q0-79 30-149t82-122t122-83t150-30q79 0 149 30t122 83t82 122t31 149zm-384 256q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100q0 53 20 99t55 82t81 55t100 20zm128 106q-61 22-123 22q-70 0-133-22v174l128-64l128 64v-174zM1920 549v1499H896v-128h896V640h-512V128H384v624q-33 8-65 20t-63 28V0h1115l549 549zm-219-37l-293-293v293h293z"/>`),
		g.Group(children),
	)
}