package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func POISolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 0q88 0 170 23t153 64t129 100t100 130t65 153t23 170q0 69-16 131t-48 125l-576 1152L448 896q-31-62-47-124t-17-132q0-88 23-170t64-153t100-129T701 88t153-65t170-23zm0 896q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100q0 53 20 99t55 82t81 55t100 20z"/>`),
		g.Group(children),
	)
}