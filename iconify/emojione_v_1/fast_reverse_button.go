package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastReverseButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#1b75bb" d="M63.793 56.913a6.877 6.877 0 0 1-6.878 6.882H6.875A6.878 6.878 0 0 1 0 56.913V6.877A6.877 6.877 0 0 1 6.875 0h50.041a6.876 6.876 0 0 1 6.878 6.877v50.036z"/><g fill="#fff"><path d="M51.938 15.04c2.305.157 4.229 1.164 5.071 2.543v27.51c-.847 1.384-2.774 2.391-5.083 2.546L29.625 31.398L51.938 15.04"/><path d="M29.1 15.04c2.304.157 4.228 1.164 5.07 2.543v27.51c-.845 1.384-2.774 2.391-5.083 2.546L6.786 31.398L29.1 15.04"/></g>`),
		g.Group(children),
	)
}