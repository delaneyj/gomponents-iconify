package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollectionSpeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M303 5q16 0 27 11.5T341 44v307q0 16-11 27t-27 11H124q-16 0-27.5-11T85 351V44q0-16 11.5-27.5T124 5h179zm-89.5 43q-17.5 0-30 12.5t-12.5 30t12.5 30t30 12.5t30-12.5t12.5-30t-12.5-30t-30-12.5zm0 288q35.5 0 60.5-25t25-60.5t-25-60.5t-60.5-25t-60.5 25t-25 60.5t25 60.5t60.5 25zM160 250.5q0-53.5 53.5-53.5t53.5 53.5t-53.5 53.5t-53.5-53.5zM43 91v341h213v43H43q-18 0-30.5-12.5T0 432V91h43z"/>`),
		g.Group(children),
	)
}