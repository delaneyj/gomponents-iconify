package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FerrySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1768 1792h152v128h-216q-12 24-29 47t-38 41t-46 29t-55 11q-41 0-71-9t-53-24t-42-30t-37-31t-39-23t-46-10q-26 0-46 9t-38 23t-36 30t-42 30t-53 23t-73 10q-42 0-72-9t-54-23t-41-30t-37-30t-38-23t-46-10q-26 0-46 9t-38 24t-37 30t-42 31t-54 23t-71 10q-31 0-56-10t-45-29t-37-41t-30-48l-2-3l-214 3v-128h152L0 1487v-157l256-85V768h152l128-256h232V256h384v256h232l128 256h152v477l256 85v157l-152 305zM896 512h128V384H896v128zm640 690V896H384v306l576-192l576 192z"/>`),
		g.Group(children),
	)
}