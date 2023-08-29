package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SketchLogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.537.827a.375.375 0 0 0-.074 0l-3.5.35a.375.375 0 0 0-.266.152L.7 5.425a.373.373 0 0 0 .012.465l6.498 7.898a.375.375 0 0 0 .58 0l6.498-7.898a.374.374 0 0 0 .087-.238v-.024a.373.373 0 0 0-.075-.203L11.303 1.33a.375.375 0 0 0-.266-.152l-3.5-.35Zm3.388 4.448v-.023l-.003.023h.003Zm.01-.1h2.253l-1.939-2.649l-.315 2.649Zm-.364-3.291l-2.527-.253l2.13 3.58l.397-3.327Zm-3.615-.253l-2.527.253l.396 3.326l2.13-3.579Zm-3.206.895L1.812 5.175h2.254L3.75 2.526ZM1.794 6.025l4.965 6.034l-2.535-5.992a.301.301 0 0 1-.015-.042H1.794Zm3.357 0L7.5 12.108l2.35-6.083h-4.7Zm5.64 0a.3.3 0 0 1-.015.042l-2.535 5.992l4.965-6.034H10.79ZM7.5 2.183l1.84 3.092H5.66L7.5 2.183Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}