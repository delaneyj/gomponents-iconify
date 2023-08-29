package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hellyhansen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.912 5.945a1.089 1.089 0 1 0-.002 2.178a1.089 1.089 0 0 0 .002-2.178zm.012.242a.85.85 0 1 1 0 1.7a.85.85 0 0 1 0-1.7zm-.332.375v.952h.18v-.352h.171l.184.352h.207l-.213-.385c.046-.017.19-.067.19-.28c0-.166-.12-.287-.323-.287h-.396zm.18.157h.167c.124 0 .184.057.184.144c0 .089-.065.143-.156.143h-.196v-.287zM0 7.039v11.016h3.684v-3.78h3.523v3.78h1.42l2.15-11.016H7.221v3.854H3.695V7.039H0zm12.127 0L9.988 18.055h3.545V14.2h3.524v3.854h3.697V7.039H17.07v3.78h-3.525v-3.78h-1.418Z"/>`),
		g.Group(children),
	)
}