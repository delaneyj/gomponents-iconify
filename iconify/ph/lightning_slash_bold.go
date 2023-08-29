package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningSlashBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M56.88 31.93a12 12 0 1 0-17.76 16.14L75.9 88.52l-36.67 39.29a12 12 0 0 0 4.56 19.43l54.44 20.41l-14 70a12 12 0 0 0 20.54 10.54l58.9-63.11l35.45 39a12 12 0 0 0 17.76-16.14Zm59.17 169l7.72-38.58a12 12 0 0 0-7.56-13.59L69 131.07l23.07-24.73l55.38 60.92Zm-4.64-132.87a12 12 0 0 1-.59-17l40.41-43.25a12 12 0 0 1 20.54 10.54l-14 70l54.44 20.41a12 12 0 0 1 4.56 19.43l-16.82 18a12 12 0 0 1-17.55-16.37l4.58-4.91l-47.19-17.69a12 12 0 0 1-7.56-13.59L140 55.07l-11.63 12.41a12 12 0 0 1-16.96.58Z"/>`),
		g.Group(children),
	)
}