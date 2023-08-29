package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardShiftUppercaseSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.606 2.191a.504.504 0 0 1 .789 0l4.598 5.802h-2.48a.502.502 0 0 0-.503.5V11H5.99V8.494a.5.5 0 0 0-.502-.5H3.007L7.606 2.19Zm1.578-.62a1.511 1.511 0 0 0-2.367 0L2.218 7.373c-.52.656-.05 1.621.789 1.621h1.978V11c0 .553.45 1.001 1.005 1.001h4.02c.556 0 1.005-.448 1.005-1.001V8.994h1.978c.84 0 1.31-.964.789-1.62L9.184 1.57ZM5.5 13a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5Z"/>`),
		g.Group(children),
	)
}