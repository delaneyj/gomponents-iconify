package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Customs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.431 2 32c0 16.568 13.432 30 30 30s30-13.432 30-30C62 15.431 48.568 2 32 2m14.709 10.471v3.801h-9.555l-.949-3.801h10.504m-9.555 4.751h9.555v4.751c0 4.764-9.555 4.764-9.555 0v-1.9h-3.801l3.801-2.851m-5.73 32.305h-14v-6h-6V29.528h6v13.999h14v6m-5.449-7.75h-5.701l13.502-14.104h14.402L33.785 42.082v-8.465l-7.81 8.16m24.601 7.752H33.785v-5.045h16.791v5.045m0-6.724H35.463l15.113-15.131v15.131"/>`),
		g.Group(children),
	)
}