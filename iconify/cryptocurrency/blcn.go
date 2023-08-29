package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blcn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM7.9 7a.9.9 0 0 0-.9.9v6.2a.9.9 0 0 0 .9.9h6.2a.9.9 0 0 0 .9-.9V7.9a.9.9 0 0 0-.9-.9zm10 0a.9.9 0 0 0-.9.9v6.2a.9.9 0 0 0 .9.9h6.2a.9.9 0 0 0 .9-.9V7.9a.9.9 0 0 0-.9-.9zm0 10a.9.9 0 0 0-.9.9v6.2a.9.9 0 0 0 .9.9h6.2a.9.9 0 0 0 .9-.9v-6.2a.9.9 0 0 0-.9-.9zm-10 0a.9.9 0 0 0-.9.9v6.2a.9.9 0 0 0 .9.9h6.2a.9.9 0 0 0 .9-.9v-6.2a.9.9 0 0 0-.9-.9z"/>`),
		g.Group(children),
	)
}