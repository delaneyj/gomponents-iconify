package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M170.5 347q-17.5 0-30-12.5T128 304t12.5-30.5t30-12.5t30 12.5T213 304t-12.5 30.5t-30 12.5zM299 155q17 0 29.5 12.5T341 197v214q0 17-12.5 29.5T299 453H43q-18 0-30.5-12.5T0 411V197q0-17 12.5-29.5T43 155h194v-43q0-27-19.5-46.5t-47-19.5T124 65.5T105 112H64q0-44 31.5-75.5T171 5t75 31.5t31 75.5v43h22zm0 256V197H43v214h256z"/>`),
		g.Group(children),
	)
}