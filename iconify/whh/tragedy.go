package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tragedy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-145-36-262-135.5T66.5 646T0 347V64Q290 0 512 0t512 64v283q0 156-66.5 299T774 888.5T512 1024zM320 192q-53 0-90.5 19T192 256.5t37.5 45T320 320t90.5-18.5t37.5-45t-37.5-45.5t-90.5-19zm-64 461q0 34 12.5 44.5T306 700t54-21t71.5-26t80.5-13t80.5 13t71.5 26t54 21t37.5-2.5T768 653q0-35-51.5-69.5t-110.5-53t-94-18.5t-94 18.5t-110.5 53T256 653zm448-461q-53 0-90.5 19T576 256.5t37.5 45T704 320t90.5-18.5t37.5-45t-37.5-45.5t-90.5-19z"/>`),
		g.Group(children),
	)
}