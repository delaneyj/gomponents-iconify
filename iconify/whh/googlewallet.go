package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Googlewallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 1024q-53 0-90.5-37.5T704 896q-9-81-40.5-169T576 576q-5 76-21.5 155.5T512 864q-24 48-58 72t-70 24q-53 0-90.5-37.5T256 832q0-60-30-132.5t-71-126T81 503q-36-15-58.5-47T0 384q0-51 31.5-89.5T96 256q35 0 74 15.5t65 43.5q14 15 30 37.5t44 64.5t43 64q0-87-8-148t-16-93t-8-48q0-53 37.5-90.5T448 64q38 0 68.5 20t46.5 53q67 69 202 282q-7-122-31-210q-30-36-30-81q0-53 37.5-90.5T832 0q33 0 55.5 16.5T928 64q4 7 10.5 20t23 57t29 92.5t23 124.5t10.5 154q0 144-18.5 249.5T960 928q-43 96-128 96z"/>`),
		g.Group(children),
	)
}