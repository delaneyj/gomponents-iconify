package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextProofingToolsGaNaDaTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M7.75 3.75a1 1 0 0 0-2 0v7.5a1 1 0 1 0 2 0v-4.5a1 1 0 0 0 0-2v-1zM23 4a1 1 0 1 0-2 0v7a1 1 0 0 0 2 0V6.75a1 1 0 0 0 0-2V4zm-8.999-1a1 1 0 0 1 1 1v.75h.25a1 1 0 1 1 0 2h-.25V11a1 1 0 0 1-2 0V8.521a1 1 0 0 1-.804.96l-2.5.5a1 1 0 0 1-1.144-1.297l1.5-4.5a1 1 0 1 1 1.897.632l-.955 2.866l.81-.162a1 1 0 0 1 1.196.961v-4.48a1 1 0 0 1 1-1zM0 4.75a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1c0 1.752-.868 4.452-3.593 5.664A.985.985 0 0 1 .1 9.906a.994.994 0 0 1 .494-1.32C1.931 7.992 2.62 6.824 2.878 5.75H1a1 1 0 0 1-1-1zM16.501 5a1 1 0 0 1 1-1h1.5a1 1 0 1 1 0 2h-.5v2.22l.758-.19a1 1 0 1 1 .485 1.94l-2 .5A1 1 0 0 1 16.5 9.5V5zM11.183 17.77l4.299-5.159a1 1 0 0 1 1.537 1.28l-5 6a1 1 0 0 1-1.476.067l-2.5-2.5a1 1 0 0 1 1.414-1.414l1.726 1.726z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}