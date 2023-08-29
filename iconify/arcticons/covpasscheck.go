package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Covpasscheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.166 28.046a5.558 5.558 0 0 1-.286 2.448q-2.265 4.783-12.896 9.461q-10.63-4.678-12.895-9.461a5.557 5.557 0 0 1-.286-2.448v-14.46Q10.74 11.67 21.196 8.88a11.325 11.325 0 0 1 5.577 0q10.456 2.789 10.456 4.183Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.141 15.528L24.086 28.311l-7.579-7.582m15.898 8.111h2.055m-1.028 2.085v-2.086m-4.806 1.564l1.78-1.028m.153 2.32l-1.043-1.806m-3.38 3.758l1.028-1.78m1.291 1.932l-1.805-1.043m-1.05 4.946v-2.057m2.085 1.028H25.6m10.607-5.02a5.65 5.65 0 0 0-8.229 6.387M32.763 5.5H42.5v9.737m0 17.526V42.5h-9.737m-17.526 0H5.5v-9.737M15.237 5.5H5.5v9.737"/>`),
		g.Group(children),
	)
}